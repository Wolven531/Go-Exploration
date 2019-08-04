package main

import "fmt"

func main() {
	// NOTE: go-lint complains type can be inferred for below line
	// var greeting string = "Hello, world!"

	// NOTE: can declare and assign using `:=` w/o var keyword
	// greeting := "Hello, "
	// NOTE: can declare and assign using `var` and `=`
	// var greeting2 = ", friend!"
	// NOTE: strings are denoted w/ double quote (") OR backtick (`)
	// NOTE: double quoted strings CANNOT contain newlines allow special escape sequences
	// name := `
	// Anthony
	// `

	// NOTE: constant value CANNOT be reassigned
	// const EXCLAMATION string = "!"
	// NOTE: constant does not need type if it can be inferred
	// const EXCLAMATION = "!"
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

	// NOTE: an unused variable causes compile error
	// NOTE: float64 is preferred floating type
	// var input float64
	// NOTE: fmt.Scanf reads user input
	// fmt.Scanf("%f", &input)

	// fmt.Printf(greeting + greeting2 + "\n")
	fmt.Println(greeting + name + greeting2)

	// NOTE: for loop in separate statements
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
