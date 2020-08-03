package main

import (
	"fmt"
	"time"
)

type Chain struct {
	Blocks []Block
}

func (c *Chain) AddBlock(b Block) {
	// Add PrevHash
	if len(c.Blocks) != 0 {
		b.PrevHash = c.Blocks[0].GetHash()
	}
	c.Blocks = append(c.Blocks, b)
}

func main() {
	chain := Chain{}

	fillChain(&chain)

	fmt.Println(chain)
	//printBlock(block, block2)
}

func fillChain(c *Chain) {
	for i := 0; i < 100; i++ {
		block := Block{Timestamp: getNowTimestamp()}
		time.Sleep(time.Millisecond * 100)
		fillBlock(&block)
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
		fmt.Printf("\n------------------- %x --------------------\n", currBlock.GetHash())
		for _, item := range currBlock.Data {
			fmt.Println(item)
		}
	}
}

func getNowTimestamp() *time.Time {
	timestamp := time.Now()
	return &timestamp
}
