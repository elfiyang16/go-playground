package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff000",
		"green": "ff001",
	}

	// var colors map[string]string // B

	// colors := make(map[string]string) //C
	colors["white"] = "ff002" // no use with struct.propertyName
	// delete(colors, "red")
	// fmt.Println(colors) //map[green:ff001 red:ff000]
	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex coder", k, "is ", v)
	}
}
