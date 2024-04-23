package go_reloaded

import (
    "testing"
)

func TestChangeArticles(t *testing.T) {
    // Test cases
    testCases := []struct {
        input    []string
        expected []string
    }{
        {input: []string{"I", "have", "a", "apple"}, expected: []string{"I", "have", "an", "apple"}},
        {input: []string{"A", "banana", "is", "tasty"}, expected: []string{"A", "banana", "is", "tasty"}},
        {input: []string{"She", "has", "an", "umbrella"}, expected: []string{"She", "has", "an", "umbrella"}},
    }

    for _, tc := range testCases {
        // Call the ChangeArticles function with the test input
        result := ChangeArticles(tc.input)
        
        // Check if the result matches the expected output
        if !equalSlices(result, tc.expected) {
            t.Errorf("ChangeArticles(%v) = %v; want %v", tc.input, result, tc.expected)
        }
    }
}

// Helper function to check if two string slices are equal
func equalSlices(slice1, slice2 []string) bool {
    if len(slice1) != len(slice2) {
        return false
    }
    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
        }
    }
    return true
}
