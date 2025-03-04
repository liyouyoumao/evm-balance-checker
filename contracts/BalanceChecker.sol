// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;


interface IERC20 {
  function totalSupply() external view returns (uint256);
  function balanceOf(address account) external view returns (uint256);
  function allowance(address owner, address spender) external view returns (uint256);
  function decimals() external view returns (uint8);
}

contract BalanceChecker {

  struct UserToken {
    address user;
    address token;
  }

  struct TokenBalance {
    UserToken userToken;
    uint256 balance;
    uint8 decimals;
    bool invalid;
  }

  /* Fallback function, don't accept any ETH */
  receive() external payable {
    revert("BalanceChecker does not accept payments");
  }

  /*
    Check the token balance and decimals of a wallet in a token contract

    Returns the balance and decimals of the token for user. Avoids possible errors:
      - return (0, 0, false) on non-erc20-contract address
  */
  function tokenBalance(address user, address token) public view returns (uint256, uint8,bool) {
    if (!IsERC20(token)) {
      return (0, 0,false);
    }
    return (IERC20(token).balanceOf(user), IERC20(token).decimals(),true);
  }

  /*
    Check the token balances and decimals of a single wallet for multiple tokens.
    Pass address(0) as a "token" address to get ETH balance.

    Returns an array of TokenBalanceInfo structs.
  */
  function balances(UserToken[] calldata userTokens) external view returns (TokenBalance[] memory) {
    TokenBalance[] memory tokenBalances = new TokenBalance[](userTokens.length);

    for (uint256 j = 0; j < userTokens.length;j++) {
      if (userTokens[j].token != address(0)) {
        (uint256 balance, uint8 decimals,bool isErc20) = tokenBalance(userTokens[j].user, userTokens[j].token);
        tokenBalances[j] = TokenBalance(userTokens[j], balance, decimals,!isErc20);
      } else {
        tokenBalances[j] = TokenBalance(userTokens[j], userTokens[j].user.balance, 18,true);
      }
    }
    return (tokenBalances);
  }


  function IsERC20(address token) public view returns (bool) {
    if (!(token.code.length > 0)) {
      return false;
    }

    try IERC20(token).totalSupply() returns (uint256) {
    } catch {
      return false;
    }

    try IERC20(token).balanceOf(address(this)) returns (uint256) {
    } catch {
      return false;
    }

    try IERC20(token).decimals() returns (uint8) {
    }catch {
      return false;
    }

    // allowance(address,address)
    try IERC20(token).allowance(address(this), address(0)) returns (uint256) {
    } catch {
      return false;
    }
    return true;
  }
}
