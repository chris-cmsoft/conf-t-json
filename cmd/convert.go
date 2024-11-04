package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"main/pkg"
	"os"
)

var ConvertCmd = &cobra.Command{
	Use:   "conf-to-json [filename]",
	Short: "conf-to-json converts conf files to JSON objects",
	Args:  cobra.ArbitraryArgs,
	Run:   convert,
}

func init() {
	ConvertCmd.Flags().StringP("file", "f", "", "input file to convert")
}

func convert(cmd *cobra.Command, args []string) {
	var scanner *bufio.Scanner
	if filename, _ := cmd.Flags().GetString("file"); filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	} else {
		file := os.Stdin
		fi, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		size := fi.Size()
		if size == 0 {
			err = errors.New("stdin is empty")
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(os.Stdin)
	}

	result, err := pkg.ConvertConfToMap(scanner)

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
