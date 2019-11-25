package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	//可能用到的函数：
	//	len(array)：获取数组长度
	//	append(array,b):将元素b添加至数组array末尾
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, block)
}

func NewGenesisBlock() *Block {
	//创世区块前置哈希为空，Data为"Genesis Block"
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	genesis := NewGenesisBlock()
	var blocks []*Block
	blocks = append(blocks, genesis)
	return &Blockchain{blocks}
}

