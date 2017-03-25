package main

import "fmt"

func main() {

	Result := veryComplexCalculation()
	fmt.Println(Result)

	String1, IntValue := retMultipleValues()
	fmt.Println(String1, IntValue)

}

func veryComplexCalculation() int {
	e := 65*3/(2+5) - 6*3
	return e
}

// we we have to return multiple parameters then
func retMultipleValues() (string, int) {
	Str := "Cool did it!!"
	Num := 10 * 5
	return Str, Num
}
