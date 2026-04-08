package operations

import "fmt"

func ExampleAddOperation() {
	add := &AddOperation{}
	result, _ := add.Execute(5, 3)
	fmt.Printf("5 + 3 = %g\n", result)
	// Output: 5 + 3 = 8
}

func ExampleSubtractOperation() {
	sub := &SubtractOperation{}
	result, _ := sub.Execute(10, 4)
	fmt.Printf("10 - 4 = %g\n", result)
	// Output: 10 - 4 = 6
}

func ExampleMultiplyOperation() {
	mul := &MultiplyOperation{}
	result, _ := mul.Execute(6, 7)
	fmt.Printf("6 * 7 = %g\n", result)
	// Output: 6 * 7 = 42
}

func ExampleDivideOperation() {
	div := &DivideOperation{}
	result, _ := div.Execute(20, 4)
	fmt.Printf("20 / 4 = %g\n", result)
	// Output: 20 / 4 = 5
}
