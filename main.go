package main

import "x-bootstrap-node/p2p"

func main() {
	p2p.InitP2P()

	// wait forever
	c := make(chan interface{})
	<-c
}
