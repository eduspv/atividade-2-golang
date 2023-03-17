package main

import "fmt"

func main() {
	var x float32
	var y float32
	var z float32

	fmt.Print("indique o seu peso em kg ")
	fmt.Scan(&x)
	fmt.Print("indique sua altura em metros ")
	fmt.Scan(&y)
	fmt.Print("o seu IMC vale ", x/(y*y))
	fmt.Scan(&z)
	if z < 18.5 {
		fmt.Print("o seu IMC esta abaixo do peso normal")

	} else if z > 18.5 && z < 25 {
		fmt.Print("o seu IMC esta ideal")

	} else {
		fmt.Print("voce esta acima do peso")
	}
}
