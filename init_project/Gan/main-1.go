package main

import (
	"fmt"
	"log"

	//"math"
	//"math/big"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	url       string
	accountID string
}

// 执行 go run *.go
func main1() {
	config := Config{
		//url: "http://localhost:7545",
		url:       "https://cloudflare-eth.com",
		accountID: "0x550FA69e0A7b61c2D3F34d4dEd7c1B3cE1327488",
	}
	// 连接Eth
	client := ConnectEth(config.url)
	// 获取账户余额
	// GetEthAccount(client, config.accountID)

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenAddress)

	address := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bal)
}
