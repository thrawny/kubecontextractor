package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	Long:  "foo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hai dear")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
