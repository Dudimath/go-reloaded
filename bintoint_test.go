package go_reloaded

import (
    "testing"
)

func TestBintoInt(t *testing.T) {
    // Test cases
    testCases := []struct {
        input    string
        expected string
    }{
        {input: "0", expected: "0"},
        {input: "1", expected: "1"},
        {input: "10", expected: "2"},
        {input: "1010", expected: "10"},
        {input: "11111111", expected: "255"},
    }

    for _, tc := range testCases {
        // Call the BintoInt function with the test input
        result := BintoInt(tc.input)
        
        // Check if the result matches the expected output
        if result != tc.expected {
            t.Errorf("BintoInt(%s) = %s; want %s", tc.input, result, tc.expected)
        }
    }
}
