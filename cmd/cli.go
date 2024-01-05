package main

import (
	"example/apisecrets/cmd/cobra"
	"fmt"
)

func main() {
	err := cobra.RootCmd.Execute()
	if err != nil {
		_ = fmt.Errorf("error in command execution: %s", err)
		return
	}
}
