package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	fmt.Print("insira o primeiro numero ")
	fmt.Scan(&n1)
	fmt.Print("insira o segundo numero ")
	fmt.Scan(&n2)
	fmt.Print("insira o terceiro numero ")
	fmt.Scan(&n3)
	if n1 < n2 && n2 < n3 {
		fmt.Print(n1, n3, n2)
	} else if n1 < n3 && n3 < n2 {
		fmt.Print(n1, n3, n2)
	} else if n2 < n1 && n1 < n3 {
		fmt.Print(n2, n1, n3)
	} else if n2 < n3 && n3 < n1 {
		fmt.Print(n2, n3, n1)
	} else if n3 < n2 && n2 < n1 {
		fmt.Print(n3, n2, n1)
	} else if n3 < n1 && n1 < n2 {
		fmt.Print(n3, n1, n2)
	}
}
