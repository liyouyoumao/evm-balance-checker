### 、ChainBase
1. 接口URL

https://api.chainbase.online/v1/account/tokens
2. 请求参数

chain_id：区块链网络ID
address：钱包地址
limit：每页返回结果数量
page：页码

```shell
curl -X GET \
"https://api.chainbase.online/v1/account/tokens?chain_id=56&address=0x868f027a5e3bd1cd29606a6681c3ddb7d3dd9b67&limit=5&page=1" \
-H "x-api-key: xxx" \
-H "accept: application/json"
```


3. 响应数据

包含token名称、符号、合约地址、余额、精度、当前美元价格、总供应量等信息。
支持分页查询，响应中包含next_page字段指示下一页页码。
数据全面，一次请求即可获取所需的所有信息。

```json
{
  "code": 0,
  "message": "ok",
  "data": [
    {
      "balance": "0x12",
      "contract_address": "0x00204f3d779c00b3187ce10ea36f3995972d79f9",
      "current_usd_price": 0,
      "decimals": 18,
      "logos": [],
      "name": "Strong",
      "symbol": "STRONG",
      "total_supply": "0",
      "urls": []
    },
    {
      "balance": "0x11b73d1a78189e25d4fe",
      "contract_address": "0x003d87d02a2a01e9e8a20f507c83e15dd83a33d1",
      "current_usd_price": 0,
      "decimals": 18,
      "logos": [],
      "name": "GT Protocol",
      "symbol": "GTAI",
      "total_supply": "75000000000000000000000000",
      "urls": []
    },
    {
      "balance": "0x1bf14b8db",
      "contract_address": "0x003f83da9868acc151be89eefa4d19838ffe5d64",
      "current_usd_price": 0,
      "decimals": 9,
      "logos": [],
      "name": "8BIT DOGE",
      "symbol": "BITD",
      "total_supply": "10000000000000000000",
      "urls": []
    },
    {
      "balance": "0xbbea97adb45072880000",
      "contract_address": "0x005614b6b3b3376ce3d50a0b665f26b63f9b7313",
      "current_usd_price": 0,
      "decimals": 18,
      "logos": [],
      "name": "FanCraze",
      "symbol": "FanCraze",
      "total_supply": "2312000000000000000000000000",
      "urls": []
    },
    {
      "balance": "0x113df6563ab71",
      "contract_address": "0x0091371eaef2da68c8dd77975358c91e8570640f",
      "current_usd_price": 0,
      "decimals": 9,
      "logos": [],
      "name": "Deflationary",
      "symbol": "DEF",
      "total_supply": "998786254207157236565193",
      "urls": []
    }
  ],
  "next_page": 2,
  "count": 2141
}
```


### 、Ankr
1. 接口URL

https://rpc.ankr.com/multichain/.../ankr_getAccountBalance=（注意：实际URL中包含特定API密钥和区块链网络参数）
2. 请求方法

POST请求，JSON RPC格式。

```shell
curl --request POST \
     --url 'https://rpc.ankr.com/multichain/xxxx/?ankr_getAccountBalance=' \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "jsonrpc": "2.0",
  "method": "ankr_getAccountBalance",
  "params": {
    "blockchain": [
      "bsc"
    ],
    "nativeFirst": true,
    "walletAddress": "0x868f027a5e3bd1cd29606a6681c3ddb7d3dd9b67"
  }
}'
```

3. 响应数据

响应包含错误信息，提示“Retry failed, reason: Failed to parse response”。
未成功获取到账户余额信息。

### 、Alchemy
1. 接口URL

https://bnb-mainnet.g.alchemy.com/v2/.../alchemy_getTokenBalances（注意：实际URL中包含特定API密钥）
2. 请求方法

POST请求，JSON RPC格式。

```shell
curl --request POST \
     --url https://bnb-mainnet.g.alchemy.com/v2/xxxx \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "id": 1,
  "jsonrpc": "2.0",
  "method": "alchemy_getTokenBalances",
  "params": [
    "0x868f027a5e3bd1cd29606a6681c3ddb7d3dd9b67"
  ]
}'
```

3. 响应数据

包含钱包地址、token合约地址及余额信息。
余额以十六进制表示，需自行转换为十进制。
精度信息需通过额外接口获取。

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "address": "0x868f027a5e3bd1cd29606a6681c3ddb7d3dd9b67",
    "tokenBalances": [
      {
        "contractAddress": "0x00204f3d779c00b3187ce10ea36f3995972d79f9",
        "tokenBalance": "0x0000000000000000000000000000000000000000000000000000000000000012"
      },
      {
        "contractAddress": "0x003d87d02a2a01e9e8a20f507c83e15dd83a33d1",
        "tokenBalance": "0x0000000000000000000000000000000000000000000011b73d1a78189e25d4fe"
      },
      {
        "contractAddress": "0x003f83da9868acc151be89eefa4d19838ffe5d64",
        "tokenBalance": "0x00000000000000000000000000000000000000000000000000000001bf4fde8a"
      },
      {
        "contractAddress": "0x005614b6b3b3376ce3d50a0b665f26b63f9b7313",
        "tokenBalance": "0x00000000000000000000000000000000000000000000bbea97adb45072880000"
      },
     ...
    ],
    "pageKey": "0x09b79fd57bd42f747eceb76beb723c827815f389"
  }
}
```

### quicknode 

未发现可以直接使用的接口，需要add-on


五、总结与建议
* Chainbase 返回的响应总量够，且单个token信息满足需要，但是免费的每秒两个请求
* Ankr 测试接口失败，pass
* Alchemy 返回最近24小时交易做多的前100token，非全部，且响应仅包含contract，balance



