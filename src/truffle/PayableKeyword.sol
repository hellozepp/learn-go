pragma solidity ^0.4.4;

contract PayableKeyword {


    // 从合约发起方向 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c 地址转入 msg.value 个以太币，单位是 wei
    function deposit() payable {

        address Account2 = 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c;
        Account2.transfer(msg.value);
    }
    //从合约发起方向某个地址转入以太币(单位是wei)，地址无效或者合约发起方余额不足时，send不会抛出异常，而是直接返回false。
//    send()方法执行时有一些风险
    //调用递归深度不能超1024。
    //如果gas不够，执行会失败。
    //所以使用这个方法要检查成功与否。
    //transfer相对send较安全
    function deposit2() payable returns (bool){

        address Account2 = 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c;
        return Account2.send(msg.value);
    }


    // 读取 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c 地址的余额
    function getAccount2Balance() constant returns (uint) {

        address Account2 = 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c;

        return Account2.balance;
    }

    // 读取合约发起方的余额
    function getOwnerBalance() constant returns (uint) {

        address Owner = msg.sender;
        return Owner.balance;
    }

}