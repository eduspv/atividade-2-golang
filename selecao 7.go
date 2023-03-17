package main

import "fmt"

func main() {
	var salario1 float32
	fmt.Print("qual é o seu salario?")
	fmt.Scan(&salario1)
	if salario1 >= 1000 {
		fmt.Print("O seu novo salario é ", salario1*1.05)

	} else {
		fmt.Print("o seu novo salario sera ", salario1*1.10)
	}
}
