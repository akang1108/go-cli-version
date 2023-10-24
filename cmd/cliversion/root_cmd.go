package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
)

var RootCmd = &cobra.Command{
	Use:     "cliversion",
	Version: version,
	Short:   "CLI auto version",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Errorf("%v", err)
		}
	},
}

func Execute() {
	RootCmd.Execute()
}
