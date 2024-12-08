package main

import ("fmt")

func main() {
	day := 5

	switch day {
	case 1,3,5 :
		fmt.Println("sattimana")
	case 2,4 :
		fmt.Println("mese")
	case 6,7 :
		fmt.Println("fino settimana")
		default :
		fmt.Println("non questa settimana")
	}
}
