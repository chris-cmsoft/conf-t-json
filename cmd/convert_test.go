package cmd

import (
	"errors"
	"reflect"
	"testing"
)

func TestConvertLineToKeyValue(t *testing.T) {
	type expectation struct {
		key   string
		value []string
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
