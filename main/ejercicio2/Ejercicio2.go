package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var numero1, numero2 int
	fmt.Println("numero1 : ")
	fmt.Scanln(&numero1)
	fmt.Println("numero2 : ")
	fmt.Scanln(&numero2)
	fmt.Println(numero1 + numero2)
}
