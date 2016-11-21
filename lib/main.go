package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost takes and ip as argument
// and listen for connections
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error:", listenErr)
		os.Exit(1)
	}

	fmt.Println("Listening on", ipAndPort)
	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", acceptErr)
	}

	fmt.Println("New connection accepted")

	for {
		handleHost(conn)
	}

}

// RunGuest takes destination ip
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal(dialErr)
	}

	for {
		handleGuest(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Print("Message received:", message)
	fmt.Print("Send Message:")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Fprint(conn, replyMessage)
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send Message:")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error:", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')

	if replyErr != nil {
		log.Fatal("Error:", replyErr)
	}
	fmt.Println("Message received:", replyMessage)
}
