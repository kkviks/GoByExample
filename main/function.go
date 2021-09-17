package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	sum := plus(1, 2)
	fmt.Println(sum)

	sum = plusPlus(1, 2, 3)
	fmt.Println(sum)
}
