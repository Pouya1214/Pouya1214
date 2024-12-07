package main

import (
	"fmt"
)

func main() {
	num := 20
	if num > 10 {
		fmt.Println("num e grane 10")
		if num > 15 {
			fmt.Println("num e grande 15")
		}
	} else {
		fmt.Println("num e grande a 10")
	}
}
