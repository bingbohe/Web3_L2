package main

import (
	"fmt"
	"log"

	// go-eth库
	"github.com/ethereum/go-ethereum/ethclient"
)

// 连接Eth虚拟机
// 连接服务 ethclient.Dial("")
func ConnectEth(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	return client
}
