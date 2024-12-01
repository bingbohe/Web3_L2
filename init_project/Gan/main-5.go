package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main5() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMBZ1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0x550FA69e0A7b61c2D3F34d4dEd7c1B3cE1327488")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("以太币余额(wei)", balance)
	fbalance(balance)

	blockNumber := big.NewInt(6609988)
	thebalance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("特定区块以太币余额(wei)", thebalance)
	fbalance(thebalance)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("处理中余额", pendingBalance)

}

// Wei转换为以太币
func fbalance(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("以太币余额(eth)", ethValue)
	return ethValue
}
