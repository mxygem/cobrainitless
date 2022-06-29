package main

import (
	"log"
	"os"

	"github.com/mxygem/cobrainitless/cmd"
)

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}

func Execute() error {
	cmd, err := cmd.AccounterCmd()
	if err != nil {
		log.Fatalf("cmd setup failed: %v\n", err)
	}

	return cmd.Execute()
}
