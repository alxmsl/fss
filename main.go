package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fss",
		Short: "File Storage Service",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
