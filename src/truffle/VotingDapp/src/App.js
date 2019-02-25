import React, { Component } from 'react'
import VotingContract from '../build/contracts/Voting.json'
import getWeb3 from './utils/getWeb3'

import './css/oswald.css'
import './css/open-sans.css'
import './css/pure-min.css'
import './App.css'


const contractAddress = "0x160555aaa125581fd6765ac57184db6a921dd12e";
var votingContractInstance;


var _modifyVotingCount = (candidates,i,votingCount) => {

    console.log("---------");
    console.log(candidates);
    console.log(i);
    console.log(votingCount);

    let obj = candidates[i];
    obj.votingCount = votingCount;
    return candidates;
}


class App extends Component {
  constructor(props) {
    super(props)

    this.state = {
      candidates: [
                    {
                      "name": "liyuechun",
                      "id": 100,
                      "votingCount": 0
                    },
                    {
                      "name": "jiangzong",
                      "id": 101,
                      "votingCount": 0
                    },
                    {
                      "name": "liudehua",
                      "id": 102,
                      "votingCount": 0
                    },
                    {
                      "name": "zhangbozhi",
                      "id": 103,
                      "votingCount": 0
                    }
                  ],
      candidatesVoteCount: ["0","0","0","0"],
      web3: null
    }
  }

  componentWillMount() {
    // Get network provider and web3 instance.
    // See utils/getWeb3 for more info.

    getWeb3
    .then(results => {
      this.setState({
        web3: results.web3
      })

      // Instantiate contract once web3 provided.
      this.instantiateContract()
    })
    .catch(() => {
      console.log('Error finding web3.')
    })
  }

  instantiateContract() {
    /*
     * SMART CONTRACT EXAMPLE
     *
     * Normally these functions would be called in the context of a
     * state management library, but for convenience I've placed them here.
     */

    const contract = require('truffle-contract')
    const votingContract = contract(VotingContract)
    votingContract.setProvider(this.state.web3.currentProvider)

    // Declaring this for later so we can chain functions on SimpleStorage.


    // Get accounts.
    this.state.web3.eth.getAccounts((error, accounts) => {
      console.log("this.state.web3.eth.getAccounts");
      votingContract.at(contractAddress).then((instance) => {
        console.log("----------");
        console.log(instance);
        votingContractInstance = instance;
        for (let i = 0; i < this.state.candidates.length; i++) {
            let object = this.state.candidates[i];
            console.log(accounts[0]);
            console.log(votingContractInstance);
            console.log(votingContractInstance.totalVotesFor(object.name));
            votingContractInstance.totalVotesFor(object.name).then(result => {
              console.log(i);
              console.log(result.c[0]);
              this.setState({
                candidates: _modifyVotingCount(this.state.candidates,i,result.c[0])
              });
            });
        }
      })
    })
  }

  render() {
    return (
      <div className="App">
      <ul>
        {
         this.state.candidates.map((object) => {
           console.log(object);
           return (

                <li key={object.id}>候选人：{object.name}          支持票数：{object.votingCount}</li>
            )
         })
        }
      </ul>

      <input
            style={{width: 200,height: 30,borderWidth: 2,marginLeft: 40}}
            placeholder="请输入候选人姓名..."
            ref="candidateInput"
      />

      <button style={{height: 30,borderWidth: 2,marginLeft: 20}} onClick={() => {
        console.log(this.refs.candidateInput);
        console.log(this.refs.candidateInput.value);
        let candidateName = this.refs.candidateInput.value;
        console.log(this.state.web3.eth.accounts[0]);
        votingContractInstance.voteForCandidate(candidateName,{from: this.state.web3.eth.accounts[0]}).then((result => {
          console.log(result);
          console.log(candidateName);
          let number = 0;
          for(let i = 0; i < this.state.candidates.length; i++) {
            let object = this.state.candidates[i];
            if (object.name === candidateName) {
              number = i;
              break;
            }
          }
          votingContractInstance.totalVotesFor(candidateName).then(result => {

            this.setState({
              candidates: _modifyVotingCount(this.state.candidates,number,result.c[0])
            });
          });

        }));
      }}>Voting</button>

      </div>
    );
  }
}

export default App
