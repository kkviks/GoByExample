package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("I don't like 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (a argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, num := range []int{7, 42} {
		r, e := f1(num)
		if e == nil {
			fmt.Println("Worked with ", r)
		} else {
			fmt.Println("Didn't work with ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
