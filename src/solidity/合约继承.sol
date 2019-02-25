pragma solidity ^ 0.4.18;

contract Animal {

    uint private _age; // public internal private

    function age() view public returns (uint) {

        return _age;
    }

    function Animal() public {
        _age = 100;
    }

}

contract Animal1 {

    uint private _age; // public internal private


    function Animal() public {
        _age = 100;
    }

}
//合约可以多继承,函数只能访问public 变量只能访问public,internal
contract Cat is Animal, Animal1 {

    function age() view public returns (uint) { //合约重写

        return 1;
    }


}
