package main

import (
	"fmt"
	"strings"
)

type SimpleUser struct {
	firstName, lastName string
}

type MyStr string

func getMethod() {
	user := &SimpleUser{firstName: "Marrie", lastName: "Ann"}
	res := user.Greeting()
	fmt.Println(res)

}

func getTypeAlias() {
	// fmt.Println((&MyStr("test")).Uppercase())
}

func (u *SimpleUser) Greeting() string {
	return fmt.Sprintf("Hello, %s %s", u.firstName, u.lastName)
}

func (s *MyStr) Uppercase() string {
	return strings.ToUpper(string(*s))
}

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func getFunc() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(compute(add))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func getClosure() {
	var x = adder()
	for i := 0; i < 3; i++ {
		fmt.Println(x(i))
	}
}
