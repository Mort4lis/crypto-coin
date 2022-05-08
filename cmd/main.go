package main

import (
	"fmt"

	"github.com/Mort4lis/crypto-coin/internal/core"
)

func main() {
	bc := core.NewBlockchain()

	bc.Add([]byte("Send 1 BTC to Ivan"))
	bc.Add([]byte("Send 2 more BTC to Ivan"))

	fmt.Println(bc)
}
