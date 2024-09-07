package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "seed"}
	rootCmd.AddCommand(createSuperUser())
	rootCmd.AddCommand(seedingUsers())
	rootCmd.Execute()
}
