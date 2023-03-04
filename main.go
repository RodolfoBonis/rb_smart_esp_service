package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	port := ":8001"

	l, err := net.Listen("tcp4", port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port", port)

	defer l.Close()

	fmt.Println("Waiting for connection")

	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println("New connection")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)

		c.Write([]byte("OK"))

		if strings.HasPrefix(temp, "umd:") {
			umd := strings.Split(temp, ":")[1]
			text := fmt.Sprintf("Umidade: %s porcento", umd)
			fmt.Println(text)
		}
	}

	c.Close()
}
