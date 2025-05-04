package main

import (
	"encoding/binary"
	"fmt"
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
		go handleServerConn(clientConn)
	}

}

func handleServerConn(clientConn net.Conn) {
	defer clientConn.Close()
	//Binary Tipe || Read
	var size uint32
	err := binary.Read(clientConn, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size) // buat buffer dengan type size
	clientConn.Read(buffer)      //Baca Clientconn dan masuk kedalam buffer
	receive := string(buffer)    // Mengubah menjadi String

	fmt.Println("Server Received : %s" + receive)

	//Write balik ke client atau membalas client
	var response string
	if receive == "I hate netvork!" {
		response = "I hate netvork toooooooooo!"
	} else {
		response = "Your Message : " + receive
	}

	binary.Write(clientConn, binary.BigEndian, uint32(len(response)))
	clientConn.Write([]byte(response))
}
