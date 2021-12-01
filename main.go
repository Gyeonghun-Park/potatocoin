package main

import (
	"github.com/Gyeonghun-Park/potatocoin/explorer"
	"github.com/Gyeonghun-Park/potatocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
