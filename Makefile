.PHONY: all clean build-contract build-contract-erc20metadata

# 默认目标
all: build-contract

# 创建输出目录
contracts/output:
	mkdir -p contracts/output

# 编译 Solidity 合约并生成 Go 绑定
build-contract: contracts/balancechecker/output
	# 编译 Solidity 合约生成 ABI 和二进制文件
	solc --abi --bin contracts/balancechecker/BalanceChecker.sol --overwrite -o contracts/balancechecker/output
	# 使用 abigen 生成 Go 绑定代码
	abigen --bin=contracts/balancechecker/output/BalanceChecker.bin --abi=contracts/balancechecker/output/BalanceChecker.abi --pkg=balancechecker --out=contracts/balancechecker/balancechecker.go

build-contract-erc20metadata: 
	# 编译 Solidity 合约生成 ABI 和二进制文件
	solc --abi --bin contracts/metadatafetcher/ERC20MetadataFetcher.sol --overwrite -o contracts/metadatafetcher/output
	# 使用 abigen 生成 Go 绑定代码
	abigen --bin=contracts/metadatafetcher/output/ERC20MetadataFetcher.bin --abi=contracts/metadatafetcher/output/ERC20MetadataFetcher.abi --pkg=metadatafetcher --out=contracts/metadatafetcher/ERC20MetadataFetcher.go


# 清理生成的文件
clean:
	rm -rf contracts/output
	rm -f contracts/balancechecker.go