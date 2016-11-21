package main

import (
	"ChatApp/lib"
	"flag"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on  the specified ip address")
	flag.Parse()

	if isHost {
		// go run main go -listen <ip>
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		// go run main go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
