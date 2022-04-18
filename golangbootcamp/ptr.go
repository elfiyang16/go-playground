package main

import "fmt"

func getPtr() {
	i, j := 1, 2
	p := &i
	fmt.Println(p)  // mem adr: 0x14000188008
	fmt.Println(*p) // 1

	*p = 3         // can't be p = 3
	fmt.Println(p) // 0x14000188008
	fmt.Println(i) //3

	p = &j
	*p = *p / 3
	fmt.Println(p) // 0x1400012c2a0
	fmt.Println(j) // 0

}
