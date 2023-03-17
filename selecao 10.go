package main

import "fmt"

func main() {
	var idade int
	fmt.Print("informe sua idade ")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Print("sua classificação é mirim")

	} else if idade >= 10 && idade <= 13 {
		fmt.Print("sua classificação é infantil")

	} else if idade >= 14 && idade <= 17 {
		fmt.Print("sua classificação é juvenil")

	} else if idade >= 18 {
		fmt.Print("sua classificação é adulto")

	}

}
