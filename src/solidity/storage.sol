pragma solidity ^ 0.4.18;


contract C {
    string public a = "liyuechun";

    function f(string memory b) pure internal {//memory(默认) 值传递 storage引用传递
        b = "13331184066";
    }

    function f1(string memory b) pure internal {//memory 值传递 storage引用传递,需要为private|internal
        byte(b)[0] = "X";
    }

    function m() view public {
        f(a);
    }

}
