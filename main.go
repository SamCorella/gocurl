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

	request := "GET /" + path + " HTTP/1.1\r\nHost: " + host + "\r\nAccept: */*\r\n\r\n"

	fmt.Println("Sending request " + request)
	connection.Write([]byte(request))
	
	buffer := make([]byte, 1024)
	bytesRead, err := connection.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer[:bytesRead]))
}