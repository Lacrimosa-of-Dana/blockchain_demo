package main


import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	//为Block生成hash，使用sha256.Sum256(data []byte)函数

	//	var buffer bytes.Buffer
	//	buffer.Write(b.PrevBlockHash)
	//	buffer.Write(int64ToBytes(b.Timestamp))
	//	buffer.Write(b.Data)
	//	hash := sha256.Sum256(buffer.Bytes())

	b.Nonce, b.Hash = NewProofOfWork(b).Run()
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}


