var contractAddress = "0x200c777c0d2e949524aecab181102c645d07ba70";
var totalSupplyHex = web3.sha3('totalSupply()').substring(0, 10);
var totalSupplyCall = { "to": contractAddress, "data": totalSupplyHex };

web3.eth.call(totalSupplyCall);
"0x000000000000000000000000000000000000000057d519abac384e1d5ea00000"

parseInt("0x000000000000000000000000000000000000000057d519abac384e1d5ea00000", 16)
2.718281828e+28
