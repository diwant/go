package coinchanger

var coinValues map[string]int

func init() {
	coinValues = map[string]int{
		"quarter": 25,
		"dime":    10,
		"nickel":  5,
		"penny":   1,
	}
}

// MakeChange ...
func MakeChange(input int) map[string]int {
	output := make(map[string]int)
	return output
}
