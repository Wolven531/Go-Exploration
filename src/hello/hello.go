package main

import (
	"fmt"
	"strconv"
)

func main() {
	// NOTE: can declare and assign using `:=` w/o var keyword
	// greeting := "Hello, "
	// NOTE: can declare and assign using `var` and `=`
	// var greeting2 = ", friend!"
	// NOTE: double quoted strings CANNOT contain newlines allow special escape sequences
	// name := `
	// Anthony
	// `

	// NOTE: multiple constant declaration
	const (
		COMMA       = ","
		EXCLAMATION = "!"
	)

	// NOTE: multiple variable declaration; no commas, just one line per declaration
	var (
		greeting  = "\tHello, "
		greeting2 = COMMA + " friend" + EXCLAMATION
		name      = `Anthony`
	)

	// NOTE: float64 is preferred floating type
	// var input float64
	// NOTE: fmt.Scanf reads user input
	// fmt.Scanf("%f", &input)

	// fmt.Printf(greeting + greeting2 + "\n")
	fmt.Println(greeting + name + greeting2)

	// NOTE: declare and initialize array of ints w/ len 10
	nums := [10]int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10, // NOTE: Go compiler requires trailing comma
	}

	// optionalSliceCapacity := 10
	// NOTE: slice can be up to len() of array but not larger
	// intSlice := make([]int, 5, optionalSliceCapacity)
	// NOTE: alternative way to make a slice
	// intSlice := nums[0:5]

	// NOTE: for is the ONLY loop structure in Go
	// for i := 1; i <= len(nums); i++ {
	// NOTE: `_` denotes unused variable in for loop
	// NOTE: `range` keyword can be used to specify which var to iterate over
	for _, i := range nums {
		description := " "
		divByThree := i%3 == 0
		divByTwo := i%2 == 0
		switch {
		case divByTwo && divByThree:
			description += "divisible by six"
		case divByThree:
			description += "divisble by three"
		case divByTwo:
			description += "even"
		default:
			description += "odd"
		}
		// if i%2 == 0 {
		// 	description += "even"
		// } else {
		// 	description += "odd"
		// }
		// NOTE: Itoa is equivalent to FormatInt(int64(i), 10)
		fmt.Println(strconv.Itoa(i) + description)
	}

	fmt.Println("Printing all numbers divisble by three from one to one hundred...")
	for a := 1; a <= 100; a++ {
		if a%3 == 0 {
			fmt.Println(a, " is divisible by three")
		} else {
			fmt.Println(a, " is NOT divisible by three")
		}
	}
}
