package main

func main() {
	println(sum(1, 2))
	println(sub(1, 2))
	println(times(1, 2))
	println(div(1, 2))
	println(mod(1, 2))
	println(sumX(1, 2, 3, 4, 5))
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
