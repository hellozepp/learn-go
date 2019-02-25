pragma solidity ^0.4.4;

contract Math {

    function m(uint a, uint b) constant returns (uint) {

        uint c = a ** b;
        //次方
        return c;
    }

    function yu() constant returns (uint) {

        uint a = 3;
        // 0b0011
        uint b = 4;
        // 0b0100

        uint c = a & b;
        // 0b0000
        return c;
        // 0
    }

    function huo() constant returns (uint) {

        uint a = 3;
        // 0b0011
        uint b = 4;
        // 0b0100

        uint c = a | b;
        // 0b0111
        return c;
        // 7
    }

    function fei() constant returns (uint8) {

        uint8 a = 3;
        // 0b00000011
        uint8 c = ~a;
        // 0b11111100
        return c;
        // 0
    }

    function yihuo() constant returns (uint) {

        uint a = 3;
        // 0b0011
        uint b = 4;
        // 0b0100

        uint c = a ^ b;
        // 0b0111
        return c;
        // 252
    }

    function leftShift() constant returns (uint8) {

        uint8 a = 8;
        // 0b00001000
        uint8 c = a << 2;
        // 0b00100000
        return c;
        // 32
    }

    function rightShift() constant returns (uint8) {

        uint8 a = 8;
        // 0b00001000
        uint8 c = a >> 2;
        // 0b00000010
        return c;
        // 2
    }
}


