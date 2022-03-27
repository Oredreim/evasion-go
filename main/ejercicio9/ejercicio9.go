package main

import "fmt"

func main() {
	var edad int
	var ano int = 2022
	fmt.Println("ingrese su edad: ")
	fmt.Scanln(&edad)

	for i := 1; i <= edad; i++ {
		edad2 := edad - i
		ayuda := ano - edad2
		println("su edad en el", ayuda, "era ", i)
	}
}
