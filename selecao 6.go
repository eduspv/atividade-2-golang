package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("primeiro numero=")
	fmt.Scan(&n1)
	fmt.Print("segundo numero=")
	fmt.Scan(&n2)
	if n1 >= 0 && n2 >= 0 {
		fmt.Print("a multiplicação dos numeros é ", n1*n2)
	} else {
		fmt.Print("a subtração dos numeros é ", n1-n2)
	}

}
