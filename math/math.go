package math

func Add(numbers ...int) int {
	var total int

	for _, num := range numbers {
		total += num
	}
	return total
}

func Subtract(numbers ...int) int {
	var total int
	for idx, num := range numbers {
		if idx == 0 {
			total = num
		} else {
			total -= num
		}
	}
	return total
}
func Divide(x int, y int) float32 {
	return float32(x / y)
}

func Multiply(numbers ...int) int {
	total := 1
	for _, num := range numbers {
		total *= num
	}
	return total
}
