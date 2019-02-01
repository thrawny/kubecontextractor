package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ContextName string

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "foo",
	Long:  "foo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ContextName)
	},
}


func Execute() {
	rootCmd.PersistentFlags().StringVarP(&ContextName, "context-name", "n", "", "The kubernetes context to copy")
	rootCmd.MarkPersistentFlagRequired("context-name")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
