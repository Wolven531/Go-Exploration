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
	intSlice := make([]int, 3, 9)
	fmt.Println("length of intSlice=", len(intSlice))
	fmt.Println("first element=", intSlice[0])
	fmt.Println("second element=", intSlice[1])
	fmt.Println("third element=", intSlice[2])

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

	divisbleMap := make(map[string]int)
	fmt.Println("Gathering all numbers divisble by three from one to one hundred...")
	for a := 1; a <= 100; a++ {
		if a%3 == 0 {
			divisbleMap["numIsDivisble"]++
		} else {
			divisbleMap["numNotDivisble"]++
		}
	}
	fmt.Println("Numbers ARE divisble by three:", divisbleMap["numIsDivisble"])
	fmt.Println("Numbers NOT divisble by three:", divisbleMap["numNotDivisble"])

	targetList := []float64{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("finding the smallest number in a list of", len(targetList), "elements...")
	fmt.Println("Smallest number =", lowest(targetList), " Largest number =", highest(targetList), " Total =", total(targetList), " Average =", average(targetList))

	// variadic function invocations
	fmt.Println("finding average of empty list w/ variadic func; average =", variadicAverage())
	fmt.Println("finding average of list w/ variadic func; average =", variadicAverage(targetList...))
	fmt.Println("finding average of 32, 49, 51, 86, 70 w/ variadic func; average =", variadicAverage(32, 49, 51, 86, 70))
}

func average(nums []float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	return total / float64(len(nums))
}

func highest(nums []float64) float64 {
	highest := nums[0]
	for _, num := range nums {
		if num > highest {
			highest = num
		}
	}
	return highest
}

func lowest(nums []float64) float64 {
	lowest := nums[0]
	for _, num := range nums {
		if num < lowest {
			lowest = num
		}
	}
	return lowest
}

func lowestAndHighest(nums []float64) (float64, float64) {
	return lowest(nums), highest(nums)
}

func total(nums []float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	return total
}

func variadicAverage(nums ...float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	return total / float64(len(nums))
}
