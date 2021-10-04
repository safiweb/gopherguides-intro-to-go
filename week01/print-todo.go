package main

import (
	"fmt"
)

func main() {

    //our variables
	var i, j, k = `"Go"`, 42, true

	//print string ("Go"), int (42), bool (true)
	fmt.Printf("Printing, %T (%v)!\n", i, i)
	fmt.Printf("Printing, %T (%v)!\n", j, j)
	fmt.Printf("Printing, %T (%v)!\n", k, k)
}
