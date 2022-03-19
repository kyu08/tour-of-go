package main

import (
	"fmt"
)

func main() {
	added := named(1, 2)
	fmt.Println(added)
}

func named(x, y int) (addedValue int) {
	addedValue = x + y
	return
}
