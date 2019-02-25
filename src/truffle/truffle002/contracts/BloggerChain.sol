pragma solidity ^ 0.4.18;
import "zeppelin-solidity/contracts/token/StandardToken.sol";
contract BloggerCoin is StandardToken {
  string public name = "BloggerCoin";
  string public symbol = "BLC";
  uint8 public decimals = 4;
  uint256 public INITIAL_SUPPLY = 666666;

  function BloggerCoin() {
    totalSupply = INITIAL_SUPPLY;
    balances[msg.sender] = INITIAL_SUPPLY;
  }
}
