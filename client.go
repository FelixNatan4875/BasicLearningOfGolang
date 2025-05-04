package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	// "time"
)

func main() {
	// Bikin while loop
	for {
		fmt.Println("Main Menu NetVork")
		fmt.Println("1. Input")
		fmt.Println("2. Exit")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			input()
		} else if choice == 2 {
			fmt.Println("Exited From Program. . .")
			break
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	var message string

	for {
		fmt.Println("Enter Message [at least 6 character and ends with '!'] :")
		scanner.Scan()
		message = scanner.Text()

		if len(message) < 6 {
			fmt.Println("Message is to short")
		} else if (message[len(message)-1]) != '!' {
			fmt.Println("Message must end with !")
		} else {
			break
		}
	}
	sendToServer(message)
}

func sendToServer(message string) {
	dial, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err)
	}

	//write
	binary.Write(dial, binary.BigEndian, uint32(len(message)))
	_, err = dial.Write([]byte(message))

	if err != nil {
		panic(err)
	}

	//Implement Deadline
	// dial.SetReadDeadline(time.Now().Add(3 * time.Second))

	//Read Server
	var size uint32
	err = binary.Read(dial, binary.BigEndian, &size)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, size) // buat buffer dengan type size
	dial.Read(buffer)            //Baca Clientconn dan masuk kedalam buffer
	receive := string(buffer)    // Mengubah menjadi String
	fmt.Println(receive)

	defer dial.Close()
}
