package main

import (
	"fmt"
	"time"
)

func getIf() string {
	ans := 42
	if ans != 42 {
		return "Wrong answer"
	} else {
		return "true"
	}
}

func getFor() {
	sum := 0
	for i := 8; i < 10; i++ {
		sum += 1
	}

	for ; sum < 20; sum++ {
		sum += 1
	}

	// while loop
	start := 1
	for start < 5 {
		start += start
	}
}

func getSwitch() {
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 3 - 2: // 1
		fmt.Println("odd")
	}

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fallthrough
		case 1, 2:
			fallthrough
		case 3:
			fmt.Println(i, 3)
			fallthrough
		case 4:
			fmt.Println(i, 4)
			break
		default:
			fmt.Println(" off the chart")
		}
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

func getDefer() string {
	v := "Get deferred!" //2
	defer fmt.Println(v)
	fmt.Println("Not defferred") //1

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	return "Return value of defferred\n" //3

}
