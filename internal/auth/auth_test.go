package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	testCases := []struct {
		expected string
		actual   string
	}{
		{"ApiKey 1234", "1234"},
		{"", ""},
	}

	for _, tc := range testCases {
		header := http.Header{}
		header.Add("Authorization", tc.expected)

		key, _ := GetAPIKey(header)

		if key != tc.actual {
			t.Fatalf("Expected[%v] is not equal to actual[%v]", key, tc.actual)
		}
	}

}
