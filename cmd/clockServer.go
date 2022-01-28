package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dmatusiewicz/concurrency/pkg/clock"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use: "clock",
	Run: clockServer,
}

func init() {
	rootCmd.AddCommand(clockCmd)
}

func clockServer(cmd *cobra.Command, args []string) {
	fmt.Println()
	port, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	clock.Run(port)
}
