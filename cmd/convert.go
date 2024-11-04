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

	result, err := convertConfToJson(scanner)

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}

func convertConfToJson(scanner *bufio.Scanner) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	for scanner.Scan() {
		lineText := strings.TrimSpace(scanner.Text())
		if lineText == "" {
			continue
		}

		key, value, err := convertLineToKeyValue(lineText)

		// If it looks like nesting, go on recursion deeper, and add to a nested key.
		if strings.HasSuffix(lineText, "{") {
			result[key], err = convertConfToJson(scanner)
			if err != nil {
				return nil, err
			}
			continue
		}

		// If it looks like recursion is ending, go back one recursion
		if strings.HasSuffix(lineText, "}") {
			return result, nil
		}

		if err != nil {
			return nil, err
		}

		if _, exists := result[key]; exists {
			result[key] = append(result[key].([]string), value...)
		} else {
			result[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func convertLineToKeyValue(line string) (key string, value []string, err error) {
	parts := strings.Split(strings.TrimSpace(line), " ")
	return parts[0], parts[1:], err
}
