pragma solidity ^ 0.4.18;

contract Animal {

  uint private _age; // public(默认) external =>用于外部 internal private=>用于内部


  function Animal() public {
    _age = 100;
  }

  function setAge(uint a) public {
    _age = a;
  }


  function _age() constant public returns (uint) {
//    return this._age; //不能用this指针访问私有成员
      return _age;
  }

}
