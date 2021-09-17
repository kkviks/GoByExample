package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend yaay!")
	default:
		fmt.Println("It's weekday oh no!!")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Before noon bro")
	default:
		fmt.Println("Afternoon bro")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Println("Bhagwaan hi jaane kya hai")
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI(1.2)
	whatAmI("Go is sooo fun")
}
