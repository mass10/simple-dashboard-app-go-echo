package main

import "fmt"
import "bytes"
import "crypto/md5"
import "encoding/json"
import "github.com/labstack/echo"
import "io/ioutil"
// import "github.com/labstack/echo-contrib/session"
// import "github.com/gorilla/sessions"
import "net/http"
import "os"
import "strings"
import "text/template"


func make_md5(s string) string {

	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func default_handler(ctx echo.Context) error {

	t, _ := template.ParseFiles("templates/index.html")
	content := make(map[string]string)
	content["url"] = "https://yukan-club.xyz/activate/08eheh392h2e9y32jhw29eyhas821h3382th"
	content["you"] = "オバマ"
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.HTML(http.StatusOK, string(buffer.Bytes()))
}

func to_user_id(unknown string) string {

	if strings.Contains(unknown, "@") {
		return make_md5(unknown)
	}
	return unknown
}

func load_user_data(unknown_id string) map[string]interface{} {

	os.Mkdir("data", 0777)

	user_id := to_user_id(unknown_id)

	// json exists?
	json_path := fmt.Sprintf("data/%s.json", user_id)
	file_content, _ := ioutil.ReadFile(json_path)
	json_text := string(file_content)

	var user_data map[string]interface{}
	err := json.Unmarshal([]byte(json_text), &user_data)
	if err != nil {
		return nil
	}
	return user_data
}

func login_handler(ctx echo.Context) error {

	email := ctx.FormValue("email")
	if email == "" {
		return ctx.Redirect(http.StatusMovedPermanently, "/")
	}
	user_id := make_md5(email)
	return ctx.Redirect(http.StatusMovedPermanently, "/" + user_id)
}

func str(unknown interface{}) string {

	if unknown == nil {
		return ""
	}
	return fmt.Sprintf("{}", unknown)
}

func dashboard_handler(ctx echo.Context) error {

	key := ctx.Param("key")
	user_data := load_user_data(key)
	t, _ := template.ParseFiles("templates/dashboard.html")
	content := make(map[string]string)
	content["email"] = str(user_data["email"])
	content["name"] = str(user_data["name"])
	content["company_name"] = str(user_data["company_name"])
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.HTML(http.StatusOK, string(buffer.Bytes()))
}

func main() {

	e := echo.New()
	// ========== routing ==========
	e.GET("/", default_handler)
	e.POST("/login", login_handler)
	e.GET("/:key", dashboard_handler)
	// e.GET("/dashboard", dashboard_handler)
	// listenning
	e.Logger.Fatal(e.Start(":8081"))
}
