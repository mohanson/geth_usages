var Web3 = require('web3');
var Tx = require('ethereumjs-tx')

var server = 'https://rinkeby.infura.io/v3/5c17ecf14e0d4756aa81b6a1154dc599'
var addressFrom = '0xeb1379888f6117386043b1e50aafa983006958d8'
var privKey = '----------------------------------------------------------------'
var addressTo = '0x64ff867048064db76f2987445cc8909267855ec8'

if (typeof web3 !== 'undefined') {
    web3 = new Web3(web3.currentProvider);
} else {
    web3 = new Web3(new Web3.providers.HttpProvider(server));
}

function sendSigned(txData, cb) {
    const privateKey = new Buffer(privKey, 'hex')
    const transaction = new Tx(txData)
    transaction.sign(privateKey)
    const serializedTx = transaction.serialize().toString('hex')
    web3.eth.sendSignedTransaction('0x' + serializedTx, cb)
}

web3.eth.getTransactionCount(addressFrom).then(txCount => {
    const txData = {
        nonce: web3.utils.toHex(txCount),
        gasLimit: web3.utils.toHex(25200),
        gasPrice: web3.utils.toHex(10e9), // 10 Gwei
        to: addressTo,
        from: addressFrom,
        value: web3.utils.toHex(web3.utils.toWei('1', 'wei'))
    }
    sendSigned(txData, function (err, result) {
        if (err) return console.log('error', err)
        console.log('sent', result)
    })
})
