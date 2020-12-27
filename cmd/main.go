package main

import "demochain/core"

func main(){
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jn")
	bc.SendData("Send 1 EOS to BB")
	bc.Print()
}