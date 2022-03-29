package main

import (
	"bufio"
	_ "bufio"
	"fmt"
	_ "fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
	"hello/main/ejercicio13/banner"
	"hello/main/ejercicio13/router"
	_ "hello/main/ejercicio13/router"
	"os"
	_ "os"
)

func Servidor(port string) {

	Server3X := echo.New()
	g := Server3X.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		if username == "ADMIN" && password == "1223" {
			return true, nil

		}
		return false, nil
	}))
	Server3X.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri},Ip=${remote_ip}, status=${status}\n",
	}))

	//Routes

	g.GET("/login", router.Loc)
	Server3X.GET("/main", router.Controller)
	Server3X.POST("/inf", router.Info)
	Server3X.POST("/main", router.Enviodatospost)
	Server3X.POST("/pic", router.Capture)

	//Start Server
	Server3X.Logger.Fatal(Server3X.Start(":" + port))

}

func main() {
	color.Blue(string(banner.Banner))
	fmt.Println("PORT: ")
	scan2 := bufio.NewScanner(os.Stdin)
	scan2.Scan()
	Servidor(scan2.Text())
}
