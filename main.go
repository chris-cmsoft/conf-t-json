package main

import (
	"fmt"
	"github.com/chris-cmsoft/conftojson/cmd"
	"os"
)

func main() {

	var rootCmd = cmd.ConvertCmd

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
