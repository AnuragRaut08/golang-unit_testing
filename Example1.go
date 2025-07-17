package unit_test

// CheckDivisibility checks if a number is divisible by 5 and returns a string
func CheckDivisibility(input int) string {
    if input%5 == 0 {
        return "FIVE"
    }
    return "NOT FIVE"
    
}
