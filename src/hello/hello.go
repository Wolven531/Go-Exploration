package main

import "fmt"

func main() {
	// // NOTE: go-lint complains type can be inferred for below line
	// // var greeting string = "Hello, world!"

	// // NOTE: can declare and assign using `:=` w/o var keyword
	// greeting := "Hello, "
	// // NOTE: can declare and assign using `var` and `=`
	// var greeting2 = ", friend!"
	// // NOTE: strings are denoted w/ double quote (") OR backtick (`)
	// name := `Anthony`

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
		greeting  = "Hello, "
		greeting2 = COMMA + " friend" + EXCLAMATION
		name      = `Anthony`
	)

	// fmt.Printf(greeting + greeting2 + "\n")
	fmt.Println(greeting + name + greeting2)
}
