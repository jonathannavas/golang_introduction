package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB Operation = iota
	DIV Operation = iota
	MUL Operation = iota
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value, ",cont:", i+1, "\n")
	}
}

func Calc(op Operation, x float64, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil

	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y musn't be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil

	default:
		return 0, errors.New("has been an error")
	}
}

// func Split(v int) (x int, y int) {
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func SumData(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MultipleOperations(op Operation, values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0, errors.New("empty values")
	}

	sum := values[0]

	for _, v := range values[1:] {

		switch op {
		case SUM:
			sum += v

		case SUB:
			sum -= v
		case DIV:
			if v == 0 {
				return 0, errors.New("y musn't be zero")
			}
			sum /= v
		case MUL:
			sum *= v

		default:
			return 0, errors.New("has been an error")
		}
	}
	return sum, nil
}
