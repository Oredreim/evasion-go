package main

import (
	_ "fmt"
	"os"
	_ "os"
)

func main() {
	file, err := os.Create("afile")
	if err != nil {
		panic(err)
	}
	data := []byte("hello world")
	file.Write(data)
	file.Close()
}
