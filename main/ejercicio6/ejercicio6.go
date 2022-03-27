package main

import "fmt"

func main() {
	var edad int

	fmt.Println("digite su edad: ")
	fmt.Scanln(&edad)

	if edad < 18 {
		println("Usted es menor de edad")
	} else {
		println("Usted es mayor de edad")
	}
}
