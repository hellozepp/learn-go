pragma solidity ^0.4.4;

//msg.sender就是当前调用方法时的发起人，一个合约部署后，通过钱包地址操作合约的人很多，但是如何正确判断谁是合约的拥有者，
//判断方式很简单，就是第一次部署合约时，谁出的gas，谁就对合约具有拥有权。
contract Test {

    address public _owner;

    uint public _number;

    function Test() {
        _owner = msg.sender;
        _number = 100;
    }

    function msgSenderAddress() constant returns (address) {
        return msg.sender;
    }

    function setNumberAdd1() {
        _number = _number + 5;
    }

    function setNumberAdd2() {
        if (_owner == msg.sender) {
            _number = _number + 10;
        }
    }

    function returnContractAddress() constant returns (address) {
        return this;
        //一个合约部署后，会有一个合约地址，这个合约地址就代表合约自己。
    }
 function getBalance(address addr) constant returns (uint){
        return addr.balance;//查看余额
    }
}

// 0xca35b7d915458ef540ade6068dfe2f44e8fa733c