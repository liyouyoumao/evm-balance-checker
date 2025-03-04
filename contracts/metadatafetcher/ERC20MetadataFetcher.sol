// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IERC20Metadata {
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
    function totalSupply() external view returns (uint256);
}

contract ERC20MetadataFetcher {
    struct ERC20Metadata {
        string symbol;
        uint8 decimals;
        uint256 totalSupply;
        bool valid;
    }

    function getTokenMetadata(address tokenAddress)
        public
        view
        returns (ERC20Metadata memory)
    {
        ERC20Metadata memory metadata;
        metadata.valid = false; // Assume valid initially

        if (!(tokenAddress.code.length > 0)) {
            return metadata;
        }

        IERC20Metadata token = IERC20Metadata(tokenAddress);

        // Check each function; if any fail, return
        try token.symbol() returns (string memory _symbol) {
            metadata.symbol = _symbol;
        } catch {
            return metadata;
        }

        try token.decimals() returns (uint8 _decimals) {
            metadata.decimals = _decimals;
        } catch {
            return metadata;
        }

        try token.totalSupply() returns (uint256 _totalSupply) {
            metadata.totalSupply = _totalSupply;
        } catch {
            return metadata;
        }
        metadata.valid = true;
        return metadata;
    }
}
