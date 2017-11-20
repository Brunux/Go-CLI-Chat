package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

/*
RunHost takes and IP address and opens a
listens for incoming messages.
*/
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenerErr := net.Listen("tcp", ipAndPort)
	if listenerErr != nil {
		log.Fatal("Error: ", listenerErr)
	}
	fmt.Println("Listing on ", ipAndPort)

	conn, connErro := listener.Accept()
	if connErro != nil {
		log.Fatal("Connection error: ", connErro)
	}
	fmt.Println("Connection receive!")

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, readerErr := reader.ReadString('\n')
	if readerErr != nil {
		log.Fatal("Reader Error: ", readerErr)
	}
	fmt.Println("Message received: ", msg)

	fmt.Print("Send replay: ")
	replayReader := bufio.NewReader(os.Stdin)
	replayMsg, replayMsgErr := replayReader.ReadString('\n')
	if replayMsgErr != nil {
		fmt.Println("Replay message error: ", replayMsgErr)
	}
	fmt.Fprint(conn, replayMsg)
}

/*
RunGuest takes an IP address and open a
connection to send messages.
*/
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, connErr := net.Dial("tcp", ipAndPort)
	if connErr != nil {
		log.Fatal("Connection error: ", connErr)
	}
	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readerErr := reader.ReadString('\n')
	if readerErr != nil {
		log.Fatal(readerErr)
	}
	fmt.Fprint(conn, message)

	replayReader := bufio.NewReader(conn)
	replayMsg, replayMsgErr := replayReader.ReadString('\n')
	if replayMsgErr != nil {
		log.Fatal("Replay Message Error: ", replayMsgErr)
	}
	fmt.Println("Message received: ", replayMsg)
}
