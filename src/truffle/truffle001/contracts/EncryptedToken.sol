pragma solidity ^ 0.4 .4;
contract EncryptedToken {
  uint256 INITIAL_SUPPLY = 666666;
  mapping(address => uint256) balances;

  function EncryptedToken() {
    balances[msg.sender] = INITIAL_SUPPLY;
  }
  //
  function transfer(address _to, uint256 _amount) {
    assert(balances[msg.sender] > _amount);
    balances[msg.sender] -= _amount;
    balances[_to] += _amount;
  }
  //
  function balanceOf(address _owner) constant returns(uint256) {
    return balances[_owner];
  }
}
