package main

import (
	"github.com/Gyeonghun-Park/potatocoin/cli"
	"github.com/Gyeonghun-Park/potatocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
