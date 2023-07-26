package main

import "fmt"

func scan() {
	var x int

	fmt.Scan(&x)
	fmt.Println("read number", x, "from stdin")

	fmt.Println(x)

}
