package go_reloaded

import (
    "testing"
    "reflect"
)

func TestPunctuations(t *testing.T) {
    // Test cases
    testCases := []struct {
        input    []string
        expected []string
    }{
        {input: []string{"I", "have", ",", "an", "apple", "."}, expected: []string{"I", "have,", "an", "apple."}},
        {input: []string{"She", "went", ",", "to", "the", "store", "!"}, expected: []string{"She", "went,", "to", "the", "store!"}},
        {input: []string{"It", "was", "a", "beautiful", "day", ":", "sunny", "and", "warm", "."}, expected: []string{"It", "was", "a", "beautiful", "day:", "sunny", "and", "warm."}},
    }

    for _, tc := range testCases {
        // Call the Punctuations function with the test input
        result := Punctuations(tc.input)
        
        // Check if the result matches the expected output
        if !reflect.DeepEqual(result, tc.expected) {
            t.Errorf("Punctuations(%v) = %v; want %v", tc.input, result, tc.expected)
        }
    }
}
