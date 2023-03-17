package main

import "fmt"

func main() {
	var n1 float32
	var n2 float32
	var n3 float32
	fmt.Print("digite o primeiro numero ")
	fmt.Scan(&n1)
	fmt.Print("digite o segundo numero ")
	fmt.Scan(&n2)
	fmt.Print("digite o terceiro numero")
	fmt.Scan(&n3)
	if n1 < n2 && n1 < n3 {
		fmt.Print("o numero menor é ", n1)

	} else if n2 < n1 && n2 < n3 {
		fmt.Print("o numero menor é ", n2)

	} else {
		fmt.Print("O numero menor é ", n3)
	}

}
