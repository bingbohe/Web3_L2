package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var contractABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":false,"internalType":"uint256","name":"ethSent","type":"uint256"},{"indexed":false,"internalType":"string","name":"message","type":"string"}],"name":"MintAttempt","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"user","type":"address"},{"indexed":false,"internalType":"uint256","name":"ethAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"tokenAmount","type":"uint256"}],"name":"MintSuccess","type":"event"}]`

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMBZ1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}
	//生成私钥
	privateKey, err := crypto.HexToECDSA("cf27fffdd1a6b9ea37a6a7757f6c5f3712a68d7560c3497575154db6e350414f")
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 获取公钥的地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fromAddress := common.HexToAddress("0x550FA69e0A7b61c2D3F34d4dEd7c1B3cE1327488")
	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 获取gasPrice
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 发送的以太币
	//value := big.NewInt(10000000000000000)
	value := big.NewInt(0)

	// 接收方地址
	toAddress := common.HexToAddress("0xa62260594EcD96512165f84f0FCc78277647511a")
	// 合约地址
	tokenAddress := common.HexToAddress("0x80b04e661c96d3c74889fc43565889bc6705c2fc")
	// 获取合约
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))
	// 准备合约参数
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("10000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	// 构造交易数据
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	/*
		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &toAddress,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(gasLimit) // 23256
	*/
	gasLimit := uint64(10000000)
	// 构建交易
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gas, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc

	// 查询合约事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tokenAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case parsedABI.Events["MintAttempt"].ID.Hex():
			event := ju. {
				User    common.Address
				EthSent *big.Int
				Message string
			}{}

			err := parsedABI.UnpackIntoInterface(&event, "MintAttempt", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("MintAttempt - User: %s, EthSent: %s, Message: %s\n",
				event.User.Hex(), event.EthSent.String(), event.Message)

		case parsedABI.Events["MintSuccess"].ID.Hex():
			event := struct {
				User        common.Address
				EthAmount   *big.Int
				TokenAmount *big.Int
			}{}
			// 解码User
			event.User = common.HexToAddress(vLog.Topics[1].Hex())
			err := parsedABI.UnpackIntoInterface(&event, "MintSuccess", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("MintSuccess - User: %s, EthAmount: %s, TokenAmount: %s\n",
				event.User.Hex(), event.EthAmount.String(), event.TokenAmount.String())
		}
	}

}
