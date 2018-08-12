package main

import "bytes"
import "github.com/labstack/echo"
import "github.com/labstack/echo-contrib/session"
import "github.com/gorilla/sessions"
import "net/http"
import "text/template"

func default_handler(ctx echo.Context) error {

	t, _ := template.ParseFiles("templates/index.html")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.String(http.StatusOK, string(buffer.Bytes()))
	// return ctx.String(http.StatusOK, "Hello, World!")
}

func dashboard_handler(ctx echo.Context) error {

	t, _ := template.ParseFiles("templates/dashboard.html")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.String(http.StatusOK, string(buffer.Bytes()))
	// return ctx.String(http.StatusOK, "Hello, World!")
}

func main() {

	e := echo.New()

	// routing
	e.GET("/", default_handler)
	e.GET("/dashboard", dashboard_handler)

	// listenning
	e.Logger.Fatal(e.Start(":8081"))
}
