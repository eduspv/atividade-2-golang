package main

import "fmt"

func main() {
	var dia int
	fmt.Print("indique o dia da semana")
	fmt.Scan(&dia)
	if dia == 1 {
		fmt.Print("hoje e domingo")

	} else if dia == 2 {
		fmt.Print("hoje é segunda")

	} else if dia == 3 {
		fmt.Print("hoje é terça")

	} else if dia == 4 {
		fmt.Print("hoje é quarta")
	} else if dia == 5 {
		fmt.Print("hoje é quinta")
	} else if dia == 6 {
		fmt.Print("hoje é sexta")
	} else if dia == 7 {
		fmt.Print("hoje é sabado")
	}

}
