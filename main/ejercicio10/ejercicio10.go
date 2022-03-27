package main

import (
	_ "fmt"
	_ "github.com/3xploit666/AesGo"
	aesgo "github.com/3xploit666/AesGo"
)

func main() {
	data := aesgo.EncryptAes("una prueba", "123456789")
	println(data)
	decrypt := aesgo.DecryptAes("U2FsdGVkX184uypmBHaU5J7y5lmnohgHaDZr7Ic1iIo=", "123456789")
	println(decrypt)
}
