package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{
		Hash:         []byte{},
		Data:         []byte(data),
		PreviousHash: previousHash,
	}

	block.DeriveHash()
	return block
}

type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, previousBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}
