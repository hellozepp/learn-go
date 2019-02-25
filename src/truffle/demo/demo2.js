contractInstance = myContract.new({
  data: bytecode,
  gas: 1000000,
  from: account1
}, function(e, contract) {
  if (!e) {
    if (!contract.address) {
      console.log("Contract transaction send: Transaction Hash: " + contract.t ransactionHash + " waiting to be mined...");
    } else {
      console.log("Contract mined! Address: " + contract.address);
      console.log(contract);
    }
  } else {
    console.log(e)
  }
})
