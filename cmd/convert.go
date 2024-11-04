package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var ConvertCmd = &cobra.Command{
	Use:   "conf-to-json [filename]",
	Short: "conf-to-json converts conf files to JSON objects",
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		_, err := os.OpenFile(args[0], os.O_RDONLY, 0)
		if os.IsNotExist(err) {
			return fmt.Errorf("specified file does not exist")
		}
		if err != nil {
			return fmt.Errorf("invalid color specified: %s", args[0])
		}
		return nil
	},
	Run: convertConfToJson,
}

func convertConfToJson(cmd *cobra.Command, args []string) {
	result := map[string][]string{}
	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
	fmt.Println(string(output))
}

func convertLineToKeyValue(line string) (key string, value []string, err error) {
	parts := strings.Split(line, " ")
	return parts[0], parts[1:], err
}
