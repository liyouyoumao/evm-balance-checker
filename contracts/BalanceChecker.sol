pragma solidity ^0.8.20;

// ERC20 contract interface
interface IToken {
  function balanceOf(address account) external view returns (uint256);
  function decimals() external view returns (uint8);
}

contract BalanceChecker {
  struct TokenBalanceInfo {
    address token;
    uint256 balance;
    uint8 decimals;
  }

  /* Fallback function, don't accept any ETH */
  receive() external payable {
    revert("BalanceChecker does not accept payments");
  }

  /*
    Check the token balance and decimals of a wallet in a token contract

    Returns the balance and decimals of the token for user. Avoids possible errors:
      - return (0, 0) on non-contract address
      - returns (0, 0) if the contract doesn't implement balanceOf or decimals
  */
  function tokenBalance(address user, address token) public view returns (uint256, uint8) {
    if (token.code.length > 0) {  // Check if it is a contract
      try IToken(token).balanceOf(user) returns (uint256 balance) {
        try IToken(token).decimals() returns (uint8 decimals) {
          return (balance, decimals);
        } catch {
          return (0, 0);
        }
      } catch {
        return (0, 0);
      }
    } else {
      return (0, 0);
    }
  }

  /*
    Check the token balances and decimals of a single wallet for multiple tokens.
    Pass address(0) as a "token" address to get ETH balance.

    Returns an array of TokenBalanceInfo structs.
  */
  function balances(address user, address[] calldata tokens) external view returns (TokenBalanceInfo[] memory) {
    TokenBalanceInfo[] memory tokenInfos = new TokenBalanceInfo[](tokens.length);

    for (uint256 j = 0; j < tokens.length; j++) {
      if (tokens[j] != address(0)) {
        (uint256 balance, uint8 decimals) = tokenBalance(user, tokens[j]);
        tokenInfos[j] = TokenBalanceInfo(tokens[j], balance, decimals);
      } else {
        tokenInfos[j] = TokenBalanceInfo(tokens[j], user.balance, 18); 
      }
    }

    return tokenInfos;
  }
}
