package calc

import "errors"

func Sum(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		return sum, errors.New("should be more than 2 numbers")
	}
	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}
