package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:          "empty html",
			inputURL:      "https://blog.boot.dev",
			inputBody:     "",
			expected:      []string{},
			errorContains: "empty HTML",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				// getting different error than expected
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v\n\n", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				// getting error when not expecting it
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v\n\n", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				// expecting an error but not getting it
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.\n\n", i, tc.name, tc.errorContains)
				return
			}

			if len(actual) != len(tc.expected) {
				// response length different than expected
				t.Errorf("Test %v - %s FAIL: expected URL length: %d, actual: %d\n\n", i, tc.name, len(tc.expected), len(actual))
				// } else {
				// 	not_same := false
				// 	for i := range tc.expected {
				// 		if actual[i] != tc.expected[i] {
				// 			not_same := true
				// 		}
				// 	}
				// 	if not_same {
				// 		// response length different than expected
				// 		t.Errorf("Test %v - %s FAIL: expected URL length: %d, actual: %d", i, tc.name, len(tc.expected), len(actual))
				// 	}
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v\n\n", i, tc.name, tc.expected, actual)
			}

		})
	}
}
