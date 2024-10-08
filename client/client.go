package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(*conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from the server:", err)
			return
		}
		fmt.Println("Message from the server: ", msg)
	}
}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin) //
	for {
		fmt.Printf("Enter text ->")         //User input prompt
		text, err := stdin.ReadString('\n') //Read user input
		if err != nil {
			fmt.Println("Error reading from the server:", err)
			return
		}
		_, err = fmt.Fprintln(*conn, text) //Send input (text) to server
		if err != nil {
			fmt.Println("Error sending to the server: ", err)
			return
		}

	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close() //closing the connection when done
	fmt.Println("Connected to the server at", *addrPtr)
	//TODO Start asynchronously reading and displaying messages
	go read(&conn)

	//TODO Start getting and sending user messages.
	write(&conn)
}
