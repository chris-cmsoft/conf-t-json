package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var ConvertCmd = &cobra.Command{
	Use:   "conf-to-json [filename]",
	Short: "conf-to-json converts conf files to JSON objects",
	Args:  cobra.ArbitraryArgs,
	Run:   convertConfToJson,
}

func init() {
	ConvertCmd.Flags().StringP("file", "f", "", "input file to convert")
}

func convertConfToJson(cmd *cobra.Command, args []string) {
	result := map[string][]string{}
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

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		key, value, err := convertLineToKeyValue(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if _, exists := result[key]; exists {
			// Do something
			result[key] = append(result[key], value...)
		} else {
			result[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}

func convertLineToKeyValue(line string) (key string, value []string, err error) {
	parts := strings.Split(line, " ")
	return parts[0], parts[1:], err
}
