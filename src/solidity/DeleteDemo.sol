pragma solidity ^0.4.0;
contract DeleteDemo {
    uint public data = 100;
    uint[] public dataArray = [uint(1),2,3];
    event e(string _str,uint _u);
    event eArr(string _str,uint[] _u);
    
    
    function doDelete(){
        uint x = data;
        e("x",x);
        delete x;
        e("after x",x);
        
        e("data",data);
        delete data;
        e("after data",data);
        
        uint[] y = dataArray;
        eArr("y",y);
        delete dataArray;
        eArr("after y",y);
    }
}