package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2727")

	if err != nil {
		panic(err)
	}
	defer listener.Close() // matiin si listenernya
	fmt.Println("Server is Listening on port 2727")

	for {
		//cara tahu ada masuk atau tidak
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleProxyConn(clientConn)
	}

}

func handleProxyConn(clientConn net.Conn) {
	defer clientConn.Close()

	serverConn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()
	go io.Copy(serverConn, clientConn)
	io.Copy(clientConn, serverConn)
}