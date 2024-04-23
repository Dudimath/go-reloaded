package go_reloaded

import (
    "testing"
    "reflect"
)

func TestApplyTransformations(t *testing.T) {
    // Test cases
    testCases := []struct {
        input    []string
        expected []string
    }{
        {input: []string{"hello", "(up)", "world"}, expected: []string{"HELLO", "world"}},
        {input: []string{"hello", "(low)", "world"}, expected: []string{"hello", "world"}},
        {input: []string{"hello", "(cap)", "world"}, expected: []string{"Hello", "world"}},
        {input: []string{"hello", "(up,", "1)", "world"}, expected: []string{"HELLO", "world"}},
        {input: []string{"hello", "(low,", "1)", "world"}, expected: []string{"hello", "world"}},
        {input: []string{"hello", "(cap,", "1)", "world"}, expected: []string{"Hello", "world"}},
    }

    for _, tc := range testCases {
        // Call the ApplyTransformations function with the test input
        result := ApplyTransformations(tc.input)
        
        // Check if the result matches the expected output
        if !reflect.DeepEqual(result, tc.expected) {
            t.Errorf("ApplyTransformations(%v) = %v; want %v", tc.input, result, tc.expected)
        }
    }
}
