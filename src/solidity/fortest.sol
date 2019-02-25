pragma solidity ^ 0.4.18;

contract C {

  uint count;


  function f() view public returns (uint){
    for(uint i = 0; i < 100000; i++){
      count += i;
    }
    return count;
  }

}
