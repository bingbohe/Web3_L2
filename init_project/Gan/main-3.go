package main

import (
	// 钱包地址
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	// 钱包库
	"github.com/ethereum/go-ethereum/crypto"
	// 加密库
	"golang.org/x/crypto/sha3"

	// 以太坊客户端
	"github.com/ethereum/go-ethereum/ethclient"
	// 以太坊交易
)

func main3() {
	fmt.Println("---------------创建钱包-----------------")
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// 获取私钥
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 打印私钥
	fmt.Println("私钥", hexutil.Encode(privateKeyBytes)[2:])
	// 获取私钥的公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 打印公钥
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥", hexutil.Encode(publicKeyBytes)[4:])
	// 获取公钥的地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 打印公钥的地址
	fmt.Println("公钥地址", address)

	// go crypto原生库手动制作公钥地址
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(publicKeyBytes))
	fmt.Println("keccak256:", hexutil.Encode(hash.Sum(nil)[12:]))

	fmt.Println("---------------进行转账-----------------")
	// 转账的以太币数量/燃气限额/燃气价格/一个自增数(nonce)，接收地址以及可选择性的添加的数据
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	// 加载私钥（准备进行交易签名，将一个私钥转化为加载到内容进行使用）
	loadPrivateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	// 生成公钥
	loadPublicKey := loadPrivateKey.Public()
	loadPublicKeyECDSA, ok := loadPublicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 转账地址
	fromAddress := crypto.PubkeyToAddress(*loadPublicKeyECDSA)
	// 自增数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// 转账金额
	value := big.NewInt(1000000000000000000)
	// 燃气限额
	gasLimit := uint64(21000)
	// 燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 接收地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 生成未签名以太坊转账事务
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 签名交易事务
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), loadPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 打印签名
	fmt.Printf("签名交易：%s\n", signedTx.Hash().Hex())
	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("交易失败")
}
