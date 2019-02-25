pragma solidity ^ 0.4.18;

contract Computer {
  uint public _a; //尺寸 15
  address _owner;

  function Computer(uint a) public{
    _a = a;
    _owner = msg.sender;
  }

  function start() constant pure public returns (string) { //表示只读
    return "开机";
  }

  function msgSender() public view returns (address) {
    return msg.sender;
  }

  function kill() public {  //不能加constant,加了为自动运行
    if(_owner == msg.sender) {
      selfdestruct(msg.sender);
    }
  }
}
