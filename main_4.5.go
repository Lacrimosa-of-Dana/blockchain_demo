package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// t := time.Now()
	var print bool
	var add bool
	var data string
	flag.BoolVar(&print, "printchain", false, "Print the whole chain.")
	flag.BoolVar(&add, "addblock", false, "Add a block to the chain.")
	flag.StringVar(&data, "data", "Default value", "The data in the block.")
	flag.Parse()
	bc := NewBlockchain()

	if add {
		bc.AddBlock(data)
	}

	if print {
		bci := bc.Iterator()
		for {
			block := bci.Next()
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			pow := NewProofOfWork(block)
			fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
			if len(block.PrevBlockHash) == 0 {
				break
			}
		}
	}
}




