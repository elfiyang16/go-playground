package main

import (
	"fmt"
)

func main() {

	var (
		name          string = "a string"
		age, location        = 24, "London"
	)
	profession := "coder"
	action := func() string {
		return "Do sth"
	}
	fmt.Printf(name, age, location, profession, action())

	//Constants can only be character, string, boolean, or numeric values and
	//cannot be declared using the := syntax.

	const ( // Capitcal letter
		StatusOK       = 200
		StatusCreated  = 201
		StatusAccepted = 202
	)
	print(StatusOK)

	name1 := "just to check"
	fmt.Printf("%v is also known as name", name1)

	region, continent := getlocation("Matt", "LA")
	fmt.Printf("%s lives in %s", region, continent)

	Types()
	getStruct()
	getCompStruct()
	getArray()
	getMultiArray()
	getRange()
	getMap()
	getFor()
	getIf()
	getSwitch()
	getMethod()
	getTypeAlias()
	getInterface()
	getGoroutine()
	getChannel()
	getBufferedChannel()
	getRangeAndCloseChannel()
	fmt.Print(getDefer())
	getPtr()
	getSlice()
	getFunc()
	getClosure()
	getGenerics()
	getSelectGoroutine()
	getSelectGoroutine2()
}

func add(x int, y int) int { // return typed
	return x + y
}

func returnInput(x int, y int) (int, int) { // return typed
	return x, y
}

func getlocation(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	region = "CHINA"
	return // If the result parameters are named,
	//a return statement without arguments returns the current values of the results
}
