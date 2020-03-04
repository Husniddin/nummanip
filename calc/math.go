package calc

import "errors"

// Add function returns sum of two integer
func Add(numbers ...int) int {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("provide more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		return nil, sum
	}
}
