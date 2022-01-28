package cmd

import (
	"github.com/dmatusiewicz/concurrency/pkg/dial"
	"github.com/spf13/cobra"
)

var dialCmd = &cobra.Command{
	Use: "dial",
	Run: dialClient,
}

func init() {
	rootCmd.AddCommand(dialCmd)
}

func dialClient(cmd *cobra.Command, args []string) {
	dial.Run(args)
}
