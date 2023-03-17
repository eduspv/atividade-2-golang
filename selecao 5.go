package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("digite um numero ")
	fmt.Scan(&n1)
	if n1%3 == 0 && n1%5 == 0 {
		fmt.Print("ele é um numero multiplo de 3 e 5!")
	} else {
		fmt.Print("ele não é um numero multiplo de 5 e 3")
	}
}
