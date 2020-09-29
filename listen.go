package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	fmt.Println("listening @ 3000")
	if err != nil {
		fmt.Println(err)
	}
	c, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(msg)
	}
}

func handleConnection() {

}