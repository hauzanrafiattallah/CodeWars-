package kata

import "strconv"

func SumMix(arr []any) int {
	sum := 0

	for _, value := range arr {
		switch v := value.(type) {
		case int:
			sum += v
		case string:
			num, _ := strconv.Atoi(v)
			sum += num

		}
	}
	return sum

}
