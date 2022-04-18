package main

import (
	"fmt"
	"time"
)

type SimpleUser2 struct {
	firstName, lastName string
}

func (u *SimpleUser2) Name() string {
	return fmt.Sprintf("%s", u.firstName)
	// *(u.firstName) 不可以

}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string {
	return c.FullName
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())

}

//======================================

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

type error interface {
	Error() string
}

func run() error { // return an object (aka MyError) that has Error() method defined
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

//===================================
func getInterface() {
	u := &SimpleUser2{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
