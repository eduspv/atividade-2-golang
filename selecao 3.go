package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("digite um numero ")
	fmt.Scan(&n1)
	if n1%2 == 0 {
		fmt.Print("o seu numero é par")

	} else {
		fmt.Print("O seu numero é impar ")
	}
}
