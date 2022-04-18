package main

import (
	"fmt"
	"strings"
)

func getArray() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [2]string{"this", "is"}
	fmt.Printf("%q\n", b) // ["this" "is"]
	fmt.Println(b)        // [this is]
	fmt.Printf("%s", b)   // [this is] the uninterpreted bytes of the string or slice

	c := [...]bool{true, false}
	fmt.Printf("%t\n", c) // [true false]
}

func getMultiArray() {
	var multi [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multi[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1) // base 10
		}
	}
	fmt.Printf("%q\n", multi)
}

func getSlice() {
	// 需要length attribute, 要么在init, 要么make
	slice := []int{2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(slice[0:0])
	fmt.Println(slice[0:1])

	slice2 := make([]int, 3)
	fmt.Println(slice2) //[0 0 0]  init with 0
	slice2[0] = 1
	fmt.Println(slice2) //[1 0 0]
	slice2 = append(slice2, 2, 3)
	fmt.Println(slice2) //[1 0 0 2 3]

	slice3 := []int{5, 6}
	slice3 = append(slice3, slice2...) // 类似 extend
	fmt.Println(slice3)                //[5 6 1 0 0 2 3]
	fmt.Println(len(slice3))           // 7

	var slice4 []int
	fmt.Println(slice4, len(slice4), cap(slice4)) // [] 0 0
	if slice4 == nil {
		fmt.Println("nil!")
	} // nil!

	//Cap and Len
	slice5 := slice
	//len=4 cap=4 [2 3 4 5]
	fmt.Printf("len=%d cap=%d %v\n", len(slice5), cap(slice5), slice5)
	slice5 = slice5[2:]
	//len=2 cap=2 [4 5]
	fmt.Printf("len=%d cap=%d %v\n", len(slice5), cap(slice5), slice5)

	//
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func getRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for _, v := range pow {
		fmt.Printf("2**%d", v)
	}

	for _, v := range pow {
		if v > 8 {
			break
		}
		fmt.Println(v)
	}

	for _, v := range pow {
		if v == 8 {
			continue
		}
		fmt.Println(v)
	}

}

type Vertex2 struct {
	Lat, Long float64
}

func getMap() {
	pat := map[string]int{
		"A": 1,
		"B": 2,
	}
	for key, value := range pat {
		fmt.Printf("%s has number %d \n", key, value)
	}

	pat2 := make(map[string]int)
	pat2["A"] = 1
	pat2["B"] = 2
	delete(pat2, "B")
	fmt.Printf("%+v\n", pat2["B"]) //0
	v, ok := pat2["B"]
	if ok == true {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("Not available")

	// If the top-level type is just a type name,
	//you can omit it from the elements of the literal.
	var m = map[string]Vertex2{
		"Google":    {Lat: 37.42202, Long: -122.08408},
		"Bell Labs": {40.68433, -74.39967},
	}
	fmt.Println(m)

}
