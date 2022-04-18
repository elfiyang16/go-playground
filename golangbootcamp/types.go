package main

import (
	"fmt"
	"math/cmplx"
)

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	u = new(Point)
	// t = new(&Point) // &Point is not a type
	r = Point{X: 1} // Y : 0
	s = Point{}     // X:0, Y:0
)

type User struct {
	Id             int
	Name, Location string
}
type Player struct {
	*User
	GameId int
}

var (
	goIsFun            = true
	maxInt  uint64     = 1<<64 - 1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func Types() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}

func conversion(i int) {
	f := float64(i)
	fmt.Println(f)
}

func getStruct() {
	fmt.Println(p, q, u, r, s) // {1 2} &{1 2} {1 0} {0 0}
}

func getCompStruct() {
	player := Player{
		&User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}
	fmt.Printf("%+v", player) //{UserIns:{Id:42 Name:Matt Location:LA} GameId:90404}%
	fmt.Println(player.Greetings())
}

func (u *User) Greetings() string { // // methods defined on a pointer are also automatically available on the value itself
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}
