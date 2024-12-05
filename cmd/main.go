package main

import (
	"fmt"
	"os"

	"github.com/kuhahalong/ddgsearch/cmd/cmds"
)

func main() {
	if err := cmds.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
