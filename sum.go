package main

import "fmt"

func main() {
	fmt.Println("Example of a sonarcloud with go")
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func mod(a int, b int) int {
	return a % b
}

func sumX(a int, b ...int) int {
	var sum int
	for _, n := range b {
		sum += n + a
	}
	return sum
}
