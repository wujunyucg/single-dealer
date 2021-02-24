package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
	"time"
)

const get_key = `
{"address":"53786ef0df57b0d082fe29972ffea659ca5f7152","crypto":{"cipher":"aes-128-ctr","ciphertext":"c96841482943357d09e1aea762349bc0ee13e0458ac33fab01f876ec4a4d0040","cipherparams":{"iv":"9cba985dac103f22b7c8e179f88d5fe9"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"85b33f9b0edccf17e7685e7b4b654b5adb07bcbeaef7aa3472021df9687bcafc"},"mac":"d8205a1eb907ac199611bbd583975cf2411032a37fa3234d91c1d143a091d59c"},"id":"bb979b17-1d32-45ca-ac09-a82ee10835a6","version":3}
`

func get(index string) string {
	conn, err := ethclient.Dial("http://10.214.242.229:18001")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(get_key), "123456")
	if err != nil {
		log.Println(auth)
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	getData, err := NewGetData(common.HexToAddress("0xC6Ba466f378D12A3a21383BcfC580ACF90B86390"), conn)
	if err != nil {
		fmt.Println("failed to deploy sendData", err, getData)
		return ""
	}
	fmt.Println("api:", getData)
	fmt.Println(time.Now())
	result, err := getData.DataMapping(nil, index)
	fmt.Println(time.Now())
	fmt.Println("result:", result, "error:", err)
	return result
	//测试查询银行名
	//fmt.Println("addr=", add.Hex(), ts.Hash().Hex())
}

func send2(index, data string) {
	conn, err := ethclient.Dial("http://10.214.242.229:18001")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(get_key), "123456")
	if err != nil {
		log.Println(auth)
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	getData, err := NewGetData(common.HexToAddress("0xC6Ba466f378D12A3a21383BcfC580ACF90B86390"), conn)
	if err != nil {
		fmt.Println("failed to deploy sendData", err, getData)
		return
	}
	fmt.Println("api:", getData)
	for true {
		result, err := getData.Send(auth, index, data)
		if err != nil {
			fmt.Println("result:", result, "error:", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	//测试查询银行名
	//fmt.Println("addr=", add.Hex(), ts.Hash().Hex())
}
