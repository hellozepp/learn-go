pragma solidity ^0.4.0;
//GoLeft == 0,GoRight == 1, GoStraight == 2, SitStill == 3
//setGoStraight方法中，我们传入的参数的值可以是0 - 3当传入的值超出这个范围时，就会中断报错
contract EnumDemo {
    enum ActionChoices {GoLeft, GoRight, GoStraight, SitStill}
    ActionChoices _choice;
    ActionChoices constant defaultChoice = ActionChoices.GoStraight;

    function setGoStraight(ActionChoices choice) public {
        _choice = choice;
    }

    function getChoice() constant public returns (ActionChoices) {
        return _choice;
    }

    function getDefaultChoice() pure public returns (uint) {
        return uint(defaultChoice);
    }
}