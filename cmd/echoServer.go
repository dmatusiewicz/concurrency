package cmd

import (
	"github.com/dmatusiewicz/concurrency/pkg/echo"
	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use: "echo",
	Run: echoServer,
}

func init() {
	rootCmd.AddCommand(echoCmd)
}

func echoServer(cmd *cobra.Command, args []string) {
	echo.Run()
}
