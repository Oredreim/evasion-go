package methods

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	Username2 = "ADMIN"
	Password2 = "1223"
)

func Sendpost(URL, MENSAJE string) error {

	datos := strings.NewReader(MENSAJE)
	res, err := http.Post(URL, "application/x-www-form-urlencoded", datos)
	userAGent := `Mozilla/5.0 (Windows NT; Windows NT 6.1; en-US) AppleWebKit/534.6 (KHTML, like Gecko) Chrome/7.0.500.0 `
	res.Header.Set("User-Agent", userAGent)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	return nil
}

func Sendpost2(URL, MENSAJE string) error {

	datos1 := strings.NewReader(MENSAJE)
	req, _ := http.NewRequest("POST", URL, datos1)
	req.SetBasicAuth(Username2, Password2)
	cli := &http.Client{}
	resp, _ := cli.Do(req)

	////////////////////////////////////////////////////////

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	return nil
}
