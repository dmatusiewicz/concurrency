package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Run: root,
}

func Execute() {
	rootCommand.Execute()
}

func root(cmd *cobra.Command, args []string) {
	cmd.Help()
}
