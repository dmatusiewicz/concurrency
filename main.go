package main

import (
	"github.com/dmatusiewicz/concurrency/cmd"
)

func main() {
	cmd.Execute()
}

// func spinner(t time.Duration) {
// 	for {
// 		for _, j := range `-\|/` {
// 			fmt.Printf("\r%c", j)
// 			time.Sleep(t)
// 		}
// 	}
// }

// func fib(f int) int {
// 	if f < 2 {
// 		return f
// 	}
// 	return fib(f-1) + fib(f-2)
// }
