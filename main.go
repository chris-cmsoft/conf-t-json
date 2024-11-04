package main

import (
	"fmt"
	"main/cmd"
	"os"
)

func main() {

	var rootCmd = cmd.ConvertCmd

	//rootCmd.AddCommand(cmd.VerifyCmd())

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
