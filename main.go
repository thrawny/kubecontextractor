package main

import (
	"github.com/thrawny/kubectl-extract/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
