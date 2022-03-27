package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var base, altura int
	fmt.Println("Dijite la base del triangulo: ")
	fmt.Scanln(&base)
	fmt.Println("Dijite la altura del triangulo: ")
	fmt.Scanln(&altura)

	area := altura * base / 2
	println(area)

}
