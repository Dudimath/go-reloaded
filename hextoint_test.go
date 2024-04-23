package go_reloaded

import (
    "testing"
)

func TestHextoInt(t *testing.T) {
    // Test cases
    testCases := []struct {
        input    string
        expected string
    }{
        {input: "1", expected: "1"},
        {input: "a", expected: "10"},
        {input: "ff", expected: "255"},
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        // Call the HextoInt function with the test input
        result := HextoInt(tc.input)
        
        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("HextoInt(%s) = %s; want %s", tc.input, result, tc.expected)
        }
    }
}
