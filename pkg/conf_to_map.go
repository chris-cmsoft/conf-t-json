package pkg

import (
	"bufio"
	"strings"
)

func ConvertConfToMap(scanner *bufio.Scanner) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	for scanner.Scan() {
		lineText := strings.TrimSpace(scanner.Text())
		if lineText == "" {
			continue
		}

		key, value, err := ConvertLineToKeyValue(lineText)

		// If it looks like nesting, go on recursion deeper, and add to a nested key.
		if strings.HasSuffix(lineText, "{") {
			result[key], err = ConvertConfToMap(scanner)
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

func ConvertLineToKeyValue(line string) (key string, value []string, err error) {
	parts := strings.Split(strings.TrimSpace(line), " ")
	return parts[0], parts[1:], err
}
