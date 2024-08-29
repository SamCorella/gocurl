package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	url := strings.Split(os.Args[1], "/")
	host := url[2]
	path := url[3]
	port := "80"

	connection, err := net.Dial("tcp", host + ":" + port)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	fmt.Println("connecting to " + host)
	fmt.Println("Sending request GET /" + path + " HTTP/1.1")
	fmt.Println("Host: " + host)
	fmt.Println("Accept: */*")
}