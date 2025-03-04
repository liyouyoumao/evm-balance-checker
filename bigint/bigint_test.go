package bigint

import (
	"database/sql/driver"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"
)

func TestFromDecimal(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Int
		wantErr bool
	}{
		{"valid number", "100", Int{big.NewInt(100)}, false},
		{"invalid number", "abc", Int{nil}, true},
		{"invalid number", "1a", Int{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromDecimal(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Scan(t *testing.T) {
	type fields struct {
		data *big.Int
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *big.Int
	}{
		{
			name:    "valid string",
			fields:  fields{nil},
			args:    args{"100"},
			wantErr: false,
			want:    big.NewInt(100),
		},
		{
			name:    "valid string with init value",
			fields:  fields{big.NewInt(1)},
			args:    args{"100"},
			wantErr: false,
			want:    big.NewInt(100),
		},
		{
			name:    "invalid string",
			fields:  fields{},
			args:    args{"abc"},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "valid bytes",
			fields:  fields{},
			args:    args{[]byte("1")},
			wantErr: false,
			want:    big.NewInt(1),
		},
		{
			name:    "invalid bytes",
			fields:  fields{},
			args:    args{[]byte("abc")},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "number type",
			fields:  fields{},
			args:    args{100},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "null type",
			fields:  fields{},
			args:    args{nil},
			wantErr: false,
			want:    new(big.Int),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Int{tt.fields.data}
			if err := i.Scan(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Int.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(i.Int, tt.want) {
				t.Errorf("Int.Scan() want %s, got %s", tt.want, i)
			}
		})
	}
}

func TestInt_String(t *testing.T) {
	type fields struct {
		data *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"not nil", fields{big.NewInt(100)}, "100"},
		{"nil", fields{}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{Int: tt.fields.data}
			if got := i.String(); got != tt.want {
				t.Errorf("Int.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	type fields struct {
		data *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "nil",
			fields:  fields{},
			want:    []byte(`{"int":null}`),
			wantErr: false,
		},
		{
			name:    "not nil",
			fields:  fields{big.NewInt(100)},
			want:    []byte(`{"int":"100"}`),
			wantErr: false,
		},
	}

	type Object struct {
		Int Int `json:"int"`
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{Int: tt.fields.data}
			got, err := json.Marshal(Object{i})
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *big.Int
	}{
		{
			name:    "nil",
			args:    args{[]byte(`{"int": null}`)},
			wantErr: false,
			want:    new(big.Int),
		},
		{
			name:    "string number",
			args:    args{[]byte(`{"int": "1000"}`)},
			wantErr: false,
			want:    big.NewInt(1000),
		},
		{
			name:    "just a number",
			args:    args{[]byte(`{"int": 1000}`)},
			wantErr: false,
			want:    big.NewInt(1000),
		},
		{
			name:    "just a number",
			args:    args{[]byte(`{"int": 10.10}`)},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "invalid string number",
			args:    args{[]byte(`{"int": "abc"}`)},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "invalid hex number",
			args:    args{[]byte(`{"int": "0x"}`)},
			wantErr: true,
			want:    nil,
		},
		{
			name:    "valid hex number",
			args:    args{[]byte(`{"int": "0x1"}`)},
			wantErr: false,
			want:    big.NewInt(1),
		},
	}

	type Object struct {
		Int Int `json:"int"`
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o Object
			err := json.Unmarshal(tt.args.text, &o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(o.Int.Int, tt.want) {
				t.Errorf("Int.UnmarshalJSON() want %s, got %s", tt.want.String(), o.Int.String())
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		x int64
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{"100", args{100}, Int{big.NewInt(100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Value(t *testing.T) {
	type fields struct {
		data *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		want    driver.Value
		wantErr bool
	}{
		{"nil", fields{}, "0", false},
		{"100", fields{big.NewInt(100)}, "100", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{Int: tt.fields.data}
			got, err := i.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int.Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Big(t *testing.T) {
	type fields struct {
		data *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   *big.Int
	}{
		{"100", fields{big.NewInt(100)}, big.NewInt(100)},
		{"nil", fields{}, big.NewInt(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{
				Int: tt.fields.data,
			}
			if got := i.Big(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int.Big() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{"1", args{[]byte{1}}, Int{big.NewInt(1)}},
		{"2", args{[]byte{2}}, Int{big.NewInt(2)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBytes(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBig(t *testing.T) {
	type args struct {
		big *big.Int
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{"test", args{big.NewInt(10)}, Int{big.NewInt(10)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBig(tt.args.big); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Readable(t *testing.T) {
	bnb, ok := new(big.Int).SetString("593445000000000000000", 10)
	if !ok {
		t.Errorf("bad bnb decimal")
		return
	}
	type fields struct {
		data *big.Int
	}
	type args struct {
		precision uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"100", fields{big.NewInt(100)}, args{0}, 100},
		{"bnb real world tx", fields{bnb}, args{18}, 593.445},
		{"nil", fields{}, args{18}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Int{
				Int: tt.fields.data,
			}
			if got := i.Readable(tt.args.precision); got != tt.want {
				t.Errorf("Int.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromReadable(t *testing.T) {
	type args struct {
		readable  float64
		precision uint8
	}
	tests := []struct {
		name    string
		args    args
		want    Int
		wantErr bool
	}{
		{"1", args{0.234, 3}, FromBig(big.NewInt(234)), false},
		{"0.0234,3", args{0.0234, 3}, Int{}, true},
		{"0.1,6", args{1, 6}, FromBig(big.NewInt(1000000)), false},
		{"0.00000002,18", args{0.00000002, 18}, FromBig(big.NewInt(20000000000)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromReadable(tt.args.readable, tt.args.precision)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromReadable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromReadable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Zero(t *testing.T) {
	type fields struct {
		Int *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "1", fields: fields{Int: new(big.Int)}, want: true},
		{name: "2", fields: fields{Int: big.NewInt(1)}, want: false},
		{name: "3", fields: fields{Int: nil}, want: true},
		{name: "4", fields: fields{Int: big.NewInt(100)}, want: false},
		{name: "5", fields: fields{Int: big.NewInt(-100)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{Int: tt.fields.Int}
			if got := i.Zero(); got != tt.want {
				t.Errorf("Int.Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Positive(t *testing.T) {
	type fields struct {
		Int *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "0", fields: fields{new(big.Int)}, want: false},
		{name: "-1", fields: fields{big.NewInt(-1)}, want: false},
		{name: "1", fields: fields{big.NewInt(1)}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Int{
				Int: tt.fields.Int,
			}
			if got := i.Positive(); got != tt.want {
				t.Errorf("Int.Positive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustDecimal(t *testing.T) {
	type args struct {
		decimal string
	}
	tests := []struct {
		name    string
		args    args
		want    Int
		success bool
	}{
		{
			name:    "success",
			args:    args{"100"},
			want:    Int{big.NewInt(100)},
			success: true,
		},
		{
			name:    "failed",
			args:    args{"abc"},
			want:    Int{},
			success: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && tt.success {
					t.Errorf("MustDecimal() wants success")
				}
			}()
			if got := MustDecimal(tt.args.decimal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustHex(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		want    Int
		success bool
	}{
		{
			name:    "with 0x prefix",
			args:    args{"0x1"},
			want:    Int{Int: big.NewInt(1)},
			success: true,
		},
		{
			name:    "without 0x prefix",
			args:    args{"abc"},
			want:    Int{Int: big.NewInt(2748)},
			success: true,
		},
		{
			name:    "invalid param",
			args:    args{"xyz"},
			want:    Int{},
			success: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && tt.success {
					t.Errorf("MustHex() wants success")
				}
			}()
			if got := MustHex(tt.args.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUint64(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want Int
	}{
		{
			name: "case 1",
			args: args{100},
			want: Int{Int: big.NewInt(100)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUint64(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Copy(t *testing.T) {
	a := New(100)
	b := a.Copy()
	a.Sub(a.Int, big.NewInt(10))

	c := a.Copy()

	if a.Cmp(big.NewInt(90)) != 0 {
		t.Error("Copy(): a should be 90")
	}

	if b.Cmp(big.NewInt(100)) != 0 {
		t.Error("Copy(): b should be 100")
	}

	if c.Cmp(big.NewInt(90)) != 0 {
		t.Error("Copy(): c should be 90")
	}
}
