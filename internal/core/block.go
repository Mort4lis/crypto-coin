package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Hash [32]byte

type Block struct {
	timestamp int64
	hash      Hash
	prevHash  Hash
	data      []byte
}

func NewBlock(data []byte, prevHash Hash) *Block {
	ts := time.Now().UnixNano()
	payload := bytes.Join([][]byte{
		[]byte(strconv.FormatInt(ts, 10)),
		prevHash[:],
		data,
	}, nil)
	hash := sha256.Sum256(payload)

	return &Block{
		timestamp: ts,
		hash:      hash,
		prevHash:  prevHash,
		data:      data,
	}
}

func (b *Block) Hash() Hash {
	return b.hash
}

func (b *Block) String() string {
	builder := &strings.Builder{}

	_, _ = fmt.Fprintf(builder, "Previous hash: %x\n", b.prevHash)
	_, _ = fmt.Fprintf(builder, "Hash: %x\n", b.hash)
	_, _ = fmt.Fprintf(builder, "Data: %s\n", b.data)

	return builder.String()
}

func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis block"), Hash{})
}
