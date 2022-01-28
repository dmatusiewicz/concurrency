package echo

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Run() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// done = make(chan bool)
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go connHandler(conn)
	}
}

// var done chan bool
var counter int

func connHandler(conn net.Conn) {
	fmt.Println("Start - New connHandler")
	defer fmt.Println("End - New connHandler")
	s := bufio.NewScanner(conn)
	for s.Scan() {

		go echo(s.Text(), conn, 3*time.Second)
		counter++
	}
	wg.Wait()
	conn.Close()
}

func echo(s string, w io.Writer, d time.Duration) {
	fmt.Println(s)
	wg.Add(1)
	s = "\t" + s + "\n"
	ss := strings.ToUpper(s)
	time.Sleep(d)
	w.Write([]byte(ss))

	sss := strings.ToLower(ss)
	time.Sleep(d)
	w.Write([]byte(sss))
	// done <- true
	defer wg.Done()
}
