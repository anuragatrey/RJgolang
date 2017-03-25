package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(value float64) (float64, error) {
	if value < 0 {
		//errors.New is used for trapping the errors
		return 0, errors.New("Math -ve number pass to square root")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
