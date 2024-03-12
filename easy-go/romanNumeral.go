package kata

func Decode(roman string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	preValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanValues[roman[i]]
		if value < preValue {
			result -= value
		} else {
			result += value
		}
		preValue = value
	}
	return result
}
