package clock

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func Run(port int) {
	p := strconv.Itoa(port)
	ln, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting server")

	for {
		fmt.Println("Accepting connections")
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue

		}

		handleConnection(c)

	}
}

func handleConnection(c net.Conn) {
	defer fmt.Println("-----")
	defer c.Close()
	connIteration := 0
	for {

		t := time.Now()
		_, err := c.Write([]byte(t.GoString() + "\n"))
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second)
		connIteration++
		fmt.Printf("connIteration: %d\n", connIteration)
	}

}
