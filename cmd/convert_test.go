package cmd

import (
	"bufio"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestConvertConfToJson(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "Simple",
			input: `
				port 22
			`,
			expected: map[string]interface{}{
				"port": []string{"22"},
			},
		},
		{
			name: "Multi Value",
			input: `
				port 22
				port 80
			`,
			expected: map[string]interface{}{
				"port": []string{"22", "80"},
			},
		},
		{
			name: "Multi Key Value",
			input: `
				port 22
				pass 80
			`,
			expected: map[string]interface{}{
				"port": []string{"22"},
				"pass": []string{"80"},
			},
		},
		{
			name: "Nested",
			input: `
				pass {
                  port 22
                }
			`,
			expected: map[string]interface{}{
				"pass": map[string]interface{}{
					"port": []string{"22"},
				},
			},
		},
		{
			name: "Nested & Unnested",
			input: `
				port 80
				pass {
                  port 22
                }
			`,
			expected: map[string]interface{}{
				"port": []string{"80"},
				"pass": map[string]interface{}{
					"port": []string{"22"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			scanner := bufio.NewScanner(reader)
			result, _ := convertConfToJson(scanner)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("result: %v, expected: %v", result, tt.expected)
			}
		})
	}
}

func TestConvertLineToKeyValue(t *testing.T) {
	type expectation struct {
		key   string
		value interface{}
		err   error
	}
	testCases := []struct {
		name     string
		line     string
		expected expectation
	}{
		{
			name: "Simple",
			line: "port 80",
			expected: expectation{
				key:   "port",
				value: []string{"80"},
				err:   nil,
			},
		},
		{
			name: "No Value",
			line: "port",
			expected: expectation{
				key:   "port",
				value: []string{},
				err:   nil,
			},
		},
		{
			name: "Empty Line",
			line: "",
			expected: expectation{
				key:   "",
				value: []string{},
				err:   nil,
			},
		},
		{
			name: "Multi Value",
			line: "port 80 8080",
			expected: expectation{
				key:   "port",
				value: []string{"80", "8080"},
				err:   nil,
			},
		},
		{
			name: "Strips outer whitespace",
			line: "                          port 80 8080                  ",
			expected: expectation{
				key:   "port",
				value: []string{"80", "8080"},
				err:   nil,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			key, value, err := convertLineToKeyValue(tt.line)
			if !errors.Is(err, tt.expected.err) {
				t.Errorf("convertLineToKeyValue() err = %v, want %v", err, tt.expected.err)
			}
			if key != tt.expected.key {
				t.Errorf("convertLineToKeyValue() key = %v, want %v", key, tt.expected.key)
			}
			if !reflect.DeepEqual(value, tt.expected.value) {
				t.Errorf("convertLineToKeyValue() value = %v, want %v", value, tt.expected.value)
			}
		})
	}
}
