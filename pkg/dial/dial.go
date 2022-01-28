package dial

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Run(ports []string) {
	// fmt.Printf("Dialing %s\n", ports)
	// for {
	// 	for _, p := range ports {
	dial("8080")
	// 	}
	// }
}

// var done = make(chan io.Reader)
// var connection *net.TCPConn

func dial(port string) {
	tcpAddr := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	}

	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	go writeToConnection(connection)
	readConnection(os.Stdout, connection)
	// go tcpStdInReader(connection, os.Stdin)
	// time.Sleep(time.Minute)
	// tcpStdOutWriter(os.Stdout, connection)
	// <-done

	// 	fmt.Println("Zamykam wpisywanie")
	// 	done <- struct{}{}

	// }(os.Stdout, conn)

	// mustCopy(conn, os.Stdin)

	// defer conn.CloseRead()
	// defer conn.CloseWrite()
	// fmt.Println("done - pizda")
	// conn.Close()

	// fmt.Println("Program finished")
	// reader := bufio.NewReader(conn)
	// line, _, _ := reader.ReadLine()
	// fmt.Printf("data from port: %s - %s\n", port, line)
	// conn.Close()
}

func readConnection(dst io.Writer, src io.Reader) {
	fmt.Println("Reader start")
	defer fmt.Println("Reader end")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func writeToConnection(c *net.TCPConn) {
	fmt.Println("Writer start")
	defer fmt.Println("Writer end")
	// wrter := bufio.NewWriter(c)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Fprintln(c, string(line))
		// fmt.Println(scanner.Text())
		// i, err := c.Write([]byte(line))
		// // wrter.Write([]byte("test"))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(i)
	}

	c.CloseWrite()
}

// func tcpStdOutWriter(dst io.Writer, src io.Reader) {
// 	log.Println("--------- Starting tcpStdOutWriter")
// 	defer log.Println("--------- Stoping tcpStdOutWriter")
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal(err)
// 	}
// 	connection.CloseRead()
// }

// func tcpStdInReader(dst io.Writer, src io.Reader) {

// 	// log.Println("--------- Starting tcpStdInReader")
// 	// defer log.Println("--------- Stoping tcpStdInReader")
// 	// if _, err := io.Copy(dst, src); err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// done <- struct{}{}
// 	// connection.CloseWrite()
// }

// func mustCopy(dst io.Writer, src io.Reader) {
// 	io.Copy(dst, src)
// 	// err != nil {
// 	// 	fmt.Println("Koncze dzialanie.")
// 	// 	log.Fatal(err)
// 	// }
// 	fmt.Println("Zamykam wpisiwanie")

// }
