package main;

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
	"time"
)
const key = `
{
    "address": "c0eeeeb71b120acd5d63affe33bf54eb583448c5",
    "crypto": {
        "cipher": "aes-128-ctr",
        "ciphertext": "139fc928fdd4ffa5622b90e9b711f55bd5ec53dc72b54bad9a8299f81b0dae04",
        "cipherparams": {
            "iv": "083f58ea5b06114b095ebfc6de3b352b"
        },
        "kdf": "scrypt",
        "kdfparams": {
            "dklen": 32,
            "n": 262144,
            "p": 1,
            "r": 8,
            "salt": "a3047c1bee8c96ba4c7016d1e770c8cbb7c6b3059d889525050f1137226e3d98"
        },
        "mac": "3d85459f67ac441653294af771b2a87b08e28a2cacf28c5a804bc5f6bccf358a"
    },
    "id": "bddc997a-a142-4d2c-818c-4f392bcddcff",
    "version": 3
}
`
func send(index , data string){
	conn, err := ethclient.Dial("http://10.214.242.228:18001")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	sendData, err := NewSendData(common.HexToAddress("0xB9b84C3429D930d8A95Bee7D8DD503Efd6Cb8B9d"), conn)
	if err != nil {
		fmt.Println("failed to deploy sendData", err, sendData)
		return
	}
	fmt.Println("api:", sendData)
	fmt.Println(time.Now())
	result, err := sendData.Send(auth, index, data)
	ctx := context.Background()
	addressAfterMined, err :=bind.WaitMined(ctx, conn, result)
	fmt.Println(addressAfterMined)
	fmt.Println(time.Now())
	fmt.Println("result:", result, "error:", err)
	//测试查询银行名
	//fmt.Println("addr=", add.Hex(), ts.Hash().Hex())
}