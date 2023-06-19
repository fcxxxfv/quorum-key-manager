# qkm local install

## docker

1. install docker

## modify .env

1. copy the .env template, and modify the env variables

## modify config

1. modify config in deps/
2. deps/config/manifests/*

## start local dev

1. make up
2. make down

### doc

1. cd doc.quorum-key-manager-main
2. npm install
3. npm start

### swagger

1. make install-swag
2. make install-swagger
3. make serve-swagger

## verify

1. create eth account

```bash
curl -X POST 'http://localhost:8082/stores/eth-accounts-test/ethereum'

{"publicKey":"0x04a1c9f86cb91a0d8cd98e9b6647a699c312e1dffb7a37c36616b10563619ca296753c4db4880aeadcc45c41b088d9efd45f4a7e4b73e59a9f18416cf011a98919","compressedPublicKey":"0x03a1c9f86cb91a0d8cd98e9b6647a699c312e1dffb7a37c36616b10563619ca296","createdAt":"2023-06-16T01:59:53.542663314Z","updatedAt":"2023-06-16T01:59:53.542663375Z","keyId":"qkm-P26Likbqe7LNkG7","address":"0xf281d091643cc6a784b6ab60ad75f0ec16c2521a","disabled":false}
```

2. send eth

```bash
cast send -r http://127.0.0.1:8546 -f 7e654d251da770a068413677967f6d3ea2fea9e4 -i --chain 888 --value 5000000000000000000 --private-key 0x56202652fdffd802b7252a456dbd8f3ecc0352bbde76c23b40afe8aebd714e2e 0xf281d091643cc6a784b6ab60ad75f0ec16c2521a
```

3. sign tx

```bash
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0xf281d091643cc6a784b6ab60ad75f0ec16c2521a","to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567", "data":"0xafed", "value": "0x12a05f200"}], "id":1}' http://localhost:8082/nodes/geth-node

{"jsonrpc":"2.0","result":"0x855402b61d687e5d86b59b3d622a0436d2ad0441272a1f05b18517d0961b429a","error":null,"id":1}
```

4. get receipt

```bash
curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0x855402b61d687e5d86b59b3d622a0436d2ad0441272a1f05b18517d0961b429a"],"id":1}' http://localhost:8082/nodes/geth-node
```
