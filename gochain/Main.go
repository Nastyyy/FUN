package main

import (
	"fmt"
	"log"
	"time"
)

type Chain struct {
	Blocks []Block
}

func (c *Chain) AddBlock(b *Block) {
	c.Blocks = append(c.Blocks, *b)
}

func (c Chain) GetLatestHash() []byte {
	if len(c.Blocks) != 0 {
		return c.Blocks[len(c.Blocks)-1].Hash
	}

	return []byte("Gensis")
}

func main() {
	chain := Chain{}

	fmt.Println("---- Creating genesis block... ----")

	genesisBlock, _ := NewBlock(getNowTimestamp(), []byte("Gensis"))
	genesisBlock.Hash = []byte("Gensis")
	chain.AddBlock(genesisBlock)

	fillChain(&chain)

	for _, b := range chain.Blocks {
		fmt.Printf("\n| PrevHash: | %x | Hash: | %x |", b.PrevHash, b.Hash)
	}
	//printBlock(block, block2)
}

func fillChain(c *Chain) {
	for i := 0; i < 10; i++ {
		block, err := NewBlock(getNowTimestamp(), c.GetLatestHash())

		if err != nil {
			log.Fatalf("Error making block: %v", err)
		}

		time.Sleep(time.Millisecond * 50)
		fillBlock(block)
		c.AddBlock(block)
	}
}

func fillBlock(b *Block) {
	for i := 0; i < 100; i++ {
		data := Transaction{Test: i, Timestamp: getNowTimestamp()}
		b.AddData(data)
	}
}

func printBlock(b ...Block) {
	for _, currBlock := range b {
		fmt.Printf("\n------------------- %x --------------------\n", currBlock.Hash)
		for _, item := range currBlock.Data {
			fmt.Println(item)
		}
	}
}

func getNowTimestamp() *time.Time {
	timestamp := time.Now()
	return &timestamp
}
