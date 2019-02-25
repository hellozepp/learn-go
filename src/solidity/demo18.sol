pragma solidity ^0.4.18;

contract CrowdFunding {

    uint c;

    function test1(uint a) pure public {

        a = 100;
    }

    function test2() constant public {

        c = 2000;
    }

    function test3() public view returns (address) {

        return msg.sender;
    }


}
