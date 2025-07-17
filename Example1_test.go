package unit_test

import "testing"

func TestCheckDivisibility(t *testing.T) {
    input := -5
    want := "FIVE"
    got := CheckDivisibility(input)

    if want != got {
        t.Errorf("Incorrect Response: got %s, want %s", got, want)
        
    }
}
