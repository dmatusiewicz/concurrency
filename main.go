package main

import (
	"github.com/dmatusiewicz/concurrency/cmd"
)

func main() {
	cmd.Execute()

}

// package main
// import (``
// 	"fmt"
// 	"log"
// 	"net"
// 	"time"
// )

// func main() {
// 	service, err := net.Listen("tcp", ":2000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer service.Close()
// 	var conn net.Conn
// 	i := 0
// 	for {
// 		i++
// 		conn, err = service.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		now := time.Now()

// 		for {
// 			fmt.Printf("%d New connection!\n", i)
// 			_, err = conn.Write([]byte(now.String()))
// 			if err != nil {
// 				fmt.Printf("Exiting %d!\n", i)
// 				conn.Close()
// 				break
// 			}
// 			time.Sleep(time.Second)

// 		}
// 	}
// }
