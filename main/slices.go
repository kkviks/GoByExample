package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(s[2])

	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	t := []int{1, 2, 3, 4, 5}
	fmt.Println(t)

	dp := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		dp[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			dp[i][j] = i + j
		}
	}
	fmt.Println(dp)

}
