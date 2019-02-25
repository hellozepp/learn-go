pragma solidity ^0.4.0;
contract StructDemo {
    struct Funder{
        address add;
        uint amount;
    }
    
    mapping (uint => Funder) public funders;
    
    uint public numFunder;
    
    event e(string _str,address _add,uint _u);
    
    
    function newFunder(address _add,uint _amount) returns (uint){
        ++numFunder;
        funders[numFunder] = Funder(_add,_amount);//可使用字典{k:v}
        e("newFunder",_add,_amount);
    }
    
    function setFunder(uint _u,uint _amount){
        Funder f = funders[_u];
        f.amount = _amount;
        e("setFunder",f.add,f.amount);
        
    }
}