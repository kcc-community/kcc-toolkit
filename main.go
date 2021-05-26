package main

import (
	"github.com/kucoin-community-chain/kcc-toolkit/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
