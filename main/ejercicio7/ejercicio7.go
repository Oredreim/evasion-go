package main

import "fmt"

func main() {
	var numero int
	fmt.Println("digite un numero")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		println("el numero es par")
	} else {
		println("el numero es impar")
	}
}
