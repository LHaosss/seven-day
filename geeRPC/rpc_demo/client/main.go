package main

import (
	"fmt"
	"geerpc/rpc_demo/server"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	// Synchronous call
	args := &server.Args{
		A: 7,
		B: 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error ", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
