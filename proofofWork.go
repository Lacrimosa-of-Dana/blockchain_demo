package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const targetBits = 20

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	hash = sha256.Sum256(pow.prepareData(nonce))
	hashInt.SetBytes(hash[:])
	for hashInt.Cmp(pow.target) == 1 {
		nonce++
		hash = sha256.Sum256(pow.prepareData(nonce))
		hashInt.SetBytes(hash[:])
	}
	fmt.Printf("\r%x", hash)
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	// the correct hash result shall lower than target
	var hashInt big.Int
	var isValid bool
	hashInt.SetBytes(pow.block.Hash)
	if hashInt.Cmp(pow.target) < 0 {
		isValid = true
	} else {
		isValid = false
	}
	return isValid
}
