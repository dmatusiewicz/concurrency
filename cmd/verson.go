package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "asd"

var versionCommand = &cobra.Command{
	Use: "version",
	Run: version,
}

func init() {
	rootCommand.AddCommand(versionCommand)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Println("Version:", Version)
}
