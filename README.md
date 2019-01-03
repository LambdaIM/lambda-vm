# lambda-vm
## Description
Lambda uses the EVM(Ethereum Virtual Machine) as a virtual machine.

## Project dependence
EVM relies on Ethereum's statedb, trie to implement contract storage and contract status updates.  
The evm dependent library is in the [lambda-libs](https://github.com/LambdaIM/lambda-libs) project.

## New Opcode
The lambda add an instruction (ORDER) to synchronize the order of the lambda store order exchange.

```
Order(bytes32 orderID, bytes32 buyerAddress, bytes32 sellerAddress, uint256 ipAddress)
```
