package main

import (
	"fmt"
	"os"

	"github.com/ericklopezdev/karasu/internal/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "karasu",
		Short:   "Karasu VCS - a minilalist version controll system written in golang",
		Version: "0.1.0",
	}

	rootCmd.AddCommand(initCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize a new karasu repository at this directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return commands.InitRepository()
		},
	}
}
