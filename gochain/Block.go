package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Timestamp *time.Time
	PrevHash  []byte
	Hash      []byte
	Data      []Data
}

func (b *Block) AddData(data Data) {
	b.Data = append(b.Data, data)
}

func (b Block) GetHash() []byte {
	hash := sha256.New()

	timestamp := b.Timestamp
	time, err := timestamp.MarshalBinary()
	if err != nil {
		fmt.Println("Failure")
	}

	hash.Write(time)
	return hash.Sum(nil)
}

type Data interface {
	ReadData() string
}

type Transaction struct {
	Test      int
	Timestamp *time.Time
}

func (t Transaction) ReadData() string {
	return string(t.Test)
}
