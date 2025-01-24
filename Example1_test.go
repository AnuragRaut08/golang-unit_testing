package unit_test

import "testing"

func TestExample(t *testing.T) {
    // Example test logic
    result := sampleFunction(5) // Call the function 
    expected := 10

    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

// Sample function for demonstration
func sampleFunction(input int) int {
    return input * 2
}
