package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("talking @ 3000")
	c, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text)
	}
}