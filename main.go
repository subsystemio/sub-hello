package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Hello int

func (t *Hello) SayHello(in *string, reply *string) error {
	*reply = *in + " world!"
	return nil
}

func main() {

	hello := new(Hello)

	server := rpc.NewServer()
	server.RegisterName("Hello", hello)
	server.HandleHTTP("/", "/debug")

	// Listen for incoming tcp packets on specified port.
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// This statement starts go's http server on
	// socket specified by l.
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := "hello"
	var reply string
	err = client.Call("Hello.SayHello", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("Conversation: %v -> %v", args, reply)
}
