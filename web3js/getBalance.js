var Web3 = require('web3');

var server = 'https://rinkeby.infura.io/v3/5c17ecf14e0d4756aa81b6a1154dc599'

if (typeof web3 !== 'undefined') {
    web3 = new Web3(web3.currentProvider);
} else {
    web3 = new Web3(new Web3.providers.HttpProvider(server));
}

web3.eth.getBalance('0xe064bdF5E3E375379735A5EA4528E6099c27513f').then(
    r => console.log(web3.utils.fromWei(r, 'ether'))
)
