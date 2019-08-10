package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
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

	// even number generation
	getNextEven := createEvenNumberGenerator()
	fmt.Println(getNextEven(), getNextEven(), getNextEven(), getNextEven(), getNextEven(), getNextEven())
	// odd number generation
	getNextOdd := createOddNumberGenerator()
	fmt.Println(getNextOdd(), getNextOdd(), getNextOdd(), getNextOdd(), getNextOdd(), getNextOdd())

	// factorial
	fmt.Println("0! =", factorial(0), "; 5! =", factorial(5), "; 10! =", factorial(10))

	// reference swap
	firstNumber := 5.0
	secondNumber := 12.0
	fmt.Println("swapping 5 and 12 using swapValues()...")
	swapValues(&firstNumber, &secondNumber)
	fmt.Println("result: firstNumber =", firstNumber, "; secondNumber =", secondNumber)
	fmt.Println("Memory address of firstNumber =", &firstNumber)

	// using structs
	firstCircle := Circle{r: 5}
	// Circle "is a" Shape, set x and y
	firstCircle.x = 3
	firstCircle.y = 3
	fmt.Println("Circle at (3,3) w/ r=5 has area=", firstCircle.area())

	fmt.Println("Starting ping n pong; Press [Enter] to terminate")
	var c = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// Shape ...
type Shape struct {
	x float64
	y float64
}

// Circle ...
type Circle struct {
	Shape
	r float64
}

// `c` can only send strings
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// `c` is bi-directional
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// `c` can only receive strings
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func average(nums []float64) float64 {
	total := 0.0

	for _, num := range nums {
		total += num
	}

	return total / float64(len(nums))
}

func createEvenNumberGenerator() func() uint {
	num := uint(0)
	return func() (returnVal uint) {
		returnVal = num
		num += 2
		return
	}
}

// TODO: optimize this is optional start values are possible
func createOddNumberGenerator() func() uint {
	num := uint(1)
	return func() (returnVal uint) {
		returnVal = num
		num += 2
		return
	}
}

func factorial(numRemaining uint) uint {
	if numRemaining == 0 {
		return 1
	}
	return numRemaining * factorial(numRemaining-1)
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

func swapValues(num1 *float64, num2 *float64) {
	tempVal := *num1
	*num1 = *num2
	*num2 = tempVal
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
