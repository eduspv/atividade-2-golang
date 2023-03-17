package main

import (
	"fmt"
)

func main() {
	var x float32
	var y float32
	fmt.Print("digite o primeiro numero ")
	fmt.Scan(&x)
	fmt.Print("digite o segundo numero ")
	fmt.Scan(&y)
	if x > y {
		fmt.Print("o numero maior é ", x)
	} else {
		fmt.Print("o numero maior é ", y)
	}
}
