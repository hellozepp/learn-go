pragma solidity ^0.4.0;
contract blockDemo {
    bytes32 public _blockHash;
    address public _coinbase = block.coinbase;
    uint public _difficulty = block.difficulty;
    uint public _gasLimit = block.gaslimit;
    uint public _number = block.number;
    bytes public _data = msg.data;
    uint public _gas = msg.gas;
    address public _sender = msg.sender;
    bytes4 public _gis = msg.sig;
    uint public _value = msg.value;
    uint public _gasPrice = tx.gasprice;
    address public _orign = tx.origin;
    
    
    function blockDemo() payable{
        //_blockHash = block.blockhash(block.number);
    }
    
    
    function getHash(uint _u){
        _blockHash = block.blockhash(_u);
    }
}