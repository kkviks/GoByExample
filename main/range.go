package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum = ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println(num, " at ", i)
		}
	}

	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, ch := range "Go is lou <3" {
		fmt.Println(i, string(ch))
	}

}
