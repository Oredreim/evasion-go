package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var nombreempleado string
	var salario int
	fmt.Println("Cual es su nombre: ")
	fmt.Scanln(&nombreempleado)
	fmt.Println("Cual es su salario base por hora: ")
	fmt.Scanln(&salario)
	dia := salario * 8
	semana := dia * 7
	mes := dia * 30

	println("salario por dia: ", dia)
	println("salario por semana: ", semana)
	println("salario por mes: ", mes)
}
