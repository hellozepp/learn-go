pragma solidity ^0.4.4;

contract MappingExample {

    // 测试账号

    // 0xca35b7d915458ef540ade6068dfe2f44e8fa733c

    // 0x14723a09acff6d2a60dcdf7aa4aff308fddc160c

    // 0x4b0897b0513fdc7c541b6d9d7e929c4e5364d2db

    mapping(address => uint)  balances;

    function update(address a, uint newBalance) public {
        balances[a] = newBalance;
    }

    // {0xca35b7d915458ef540ade6068dfe2f44e8fa733c: 100,0x14723a09acff6d2a60dcdf7aa4aff308fddc160c: 200,0x4b0897b0513fdc7c541b6d9d7e929c4e5364d2db: 300 }

    function searchBalance(address a) constant public returns (uint) {

        return balances[a];
    }
}