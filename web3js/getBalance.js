var Web3 = require('web3');

if (typeof web3 !== 'undefined') {
    web3 = new Web3(web3.currentProvider);
} else {
    web3 = new Web3(new Web3.providers.HttpProvider('https://ropsten.infura.io/abAWsazRr6zO8zmW8J4i'));
}

web3.eth.getBalance('0xe064bdF5E3E375379735A5EA4528E6099c27513f').then(
    r => console.log(web3.utils.fromWei(r, 'ether'))
)
