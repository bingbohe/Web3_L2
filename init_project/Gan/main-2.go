package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main2() {
	fmt.Println("---------------连接信息-----------------")
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMBZ1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------查链信息-----------------")
	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------查区块信息-----------------")
	// 这两句说明指定了具体区块，如果查询最新区块，blockNumber可设置为nil
	// 获取区块高度
	blockNumber := big.NewInt(5671744)
	// 获取区块头
	header, err := client.HeaderByNumber(context.Background(), blockNumber)

	// 打印区块头信息
	fmt.Println(header.Number.Uint64())
	fmt.Println(header.Time)
	fmt.Println(header.Difficulty.Uint64())
	fmt.Println(header.Hash().Hex())

	if err != nil {
		log.Fatal(err)
	}
	// 获取区块信息（包含区块头）
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty())
	fmt.Println(block.Hash().Hex())
	fmt.Println("交易笔数:", len(block.Transactions()))

	/* 获取指定区块的交易笔数，
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70
	*/
	fmt.Println("---------------查交易信息-----------------")
	// 查询区块中交易信息
	for _, tx := range block.Transactions() {
		// 当前的交易哈希
		fmt.Println(tx.Hash().Hex())
		// 当前交易的转账金额
		fmt.Println(tx.Value().String())
		// 转账限额（指的是 完成交易的需要消耗的最大计算资源费用，既gas消耗量）
		fmt.Println(tx.Gas())
		// 交易价格（gas消耗 * gas价格）
		fmt.Println(tx.GasPrice().Uint64())
		// 交易难度
		fmt.Println(tx.Nonce())
		// 如果是合约，可以获取合约签名及如入参数据
		fmt.Println(tx.Data())
		// 交易接收方地址
		fmt.Println(tx.To().Hex())
		fmt.Println(tx.Value().String())
		// 从交易中获取交易发起者地址（NewEIP155Signer签名交易）
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println(sender.Hex())
		}

		if err != nil {
			log.Fatal(err)
		}
		break
	}
	// 按哈希值查询特定区块
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	// 按哈希值查询特定区块、交易 通过has值
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false

	fmt.Println("---------------查区块收据信息-----------------")

	// 使用区块哈希查询区块收据
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	// 使用区块高度查询区块收据
	//receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))

	// fmt.Println(*(receiptByHash[0]) == *(receiptsByNum[0]))

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.GasUsed)
		fmt.Println(receipt.Logs)
		break
	}

	txHashRecipt := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHashRecipt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	// 交易哈希
	fmt.Println(receipt.TxHash.Hex())
	// 消耗gas数量
	fmt.Println(receipt.GasUsed)
	// 交易事件
	fmt.Println(receipt.Logs)
	// 交易索引
	fmt.Println(receipt.TransactionIndex)
	// 合约地址
	fmt.Println(receipt.ContractAddress.Hex())
}
