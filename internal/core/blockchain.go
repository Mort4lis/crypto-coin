package core

import "strings"

type Blockchain struct {
	hashes []Hash
	blocks map[Hash]*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{blocks: make(map[Hash]*Block)}
	genesisBlock := NewGenesisBlock()

	bc.addBlock(genesisBlock)

	return bc
}

func (bc *Blockchain) Add(data []byte) {
	lastHash := bc.hashes[len(bc.hashes)-1]
	block := NewBlock(data, lastHash)

	bc.addBlock(block)
}

func (bc *Blockchain) addBlock(block *Block) {
	bc.blocks[block.Hash()] = block
	bc.hashes = append(bc.hashes, block.Hash())
}

func (bc *Blockchain) String() string {
	builder := strings.Builder{}

	for _, hash := range bc.hashes {
		block := bc.blocks[hash]

		builder.WriteString(block.String())
		builder.WriteByte('\n')
	}

	return builder.String()
}
