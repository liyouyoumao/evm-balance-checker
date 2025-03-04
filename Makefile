.PHONY: all clean build-contract

# 默认目标
all: build-contract

# 创建输出目录
contracts/output:
	mkdir -p contracts/output

# 编译 Solidity 合约并生成 Go 绑定
build-contract: contracts/output
	# 编译 Solidity 合约生成 ABI 和二进制文件
	solc --abi --bin contracts/BalanceChecker.sol --overwrite -o contracts/output
	# 使用 abigen 生成 Go 绑定代码
	abigen --bin=contracts/output/BalanceChecker.bin --abi=contracts/output/BalanceChecker.abi --pkg=balancechecker --out=contracts/balancechecker.go

# 清理生成的文件
clean:
	rm -rf contracts/output
	rm -f contracts/balancechecker.go