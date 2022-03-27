package main

import "fmt"

func main() {
	var nota1, nota2, nota3 float32
	fmt.Println("ingrese nota primera nota: ")
	fmt.Scanln(&nota1)
	fmt.Println("ingrese nota segunda nota: ")
	fmt.Scanln(&nota2)
	fmt.Println("ingrese nota tercera nota: ")
	fmt.Scanln(&nota3)

	notafinal := (nota1 * 0.20) + (nota2 * 0.30) + (nota3 * 0.50)
	println(notafinal)

}
