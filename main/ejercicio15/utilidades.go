package utilidades

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"net/http"
)

func Base64Decode(str string) (string, string) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {

		return "", ""
	}
	return string(data), ""
}
func Bajarcadena(code string) string {
	link, err := http.Get(code)
	userAGent := `Mozilla/5.0 (Windows NT; Windows NT 6.1; en-US) AppleWebKit/534.6 (KHTML, like Gecko) Chrome/7.0.500.0 `
	link.Header.Set("User-Agent", userAGent)

	if err != nil {
		log.Fatal(err)
	}
	reqBody, err := ioutil.ReadAll(link.Body)
	if err != nil {
		log.Fatal(err)
	}
	//println(string(reqBody))
	return string(reqBody)
}

func CreateREGG(value string) {

	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Wow6432Node\`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	if err := k.SetStringValue("SystemUpdate", value); err != nil {
		log.Fatal(err)
	}

}

func Registroleer() (string, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Wow6432Node\`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	version, _, err := k.GetStringValue("SystemUpdate")
	fmt.Println(version)
	return version, err

}
