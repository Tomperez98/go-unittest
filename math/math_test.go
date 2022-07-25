package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type TestData struct {
		Numbers []int
		Result  int
	}

	bunchOfTests := []TestData{
		{[]int{1, 3}, 4},
		{[]int{2, 3}, 5},
	}
	for _, data := range bunchOfTests {
		result := Add(data.Numbers...)
		if result != data.Result {
			panic("Bad")
		}
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(7, 3)
	if result != 4 {
		panic(result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(9, 3)
	if result != float32(3) {
		panic("Bad")
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(10, 3)
	if result != 30 {
		panic("Bad")
	}

}
