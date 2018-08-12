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

func load_user_data(unknown_id string) map[string]string {

	os.Mkdir("data", 0777)

	user_id := to_user_id(unknown_id)

	// json exists?
	json_path := fmt.Sprintf("data/%s.json", user_id)
	file_content, _ := ioutil.ReadFile(json_path)
	json_text := string(file_content)

	user_data := make(map[string]string)
	err := json.Unmarshal([]byte(json_text), &user_data)
	if err != nil {
		return user_data
	}
	return user_data
}

func store_user_data(unknown_id string, user_data map[string]string) {

	user_id := to_user_id(unknown_id)

	// store
	os.Mkdir("data", 0777)
	json_path := fmt.Sprintf("data/%s.json", user_id)
	// json_text := string(file_content)

	// var raw_data []byte
	raw_data, err := json.Marshal(&user_data)
	if err != nil {
		return
	}
	ioutil.WriteFile(json_path, raw_data, 0777)
}

func login_handler(ctx echo.Context) error {

	email := ctx.FormValue("email")
	if email == "" {
		return ctx.Redirect(http.StatusMovedPermanently, "/")
	}
	user_id := make_md5(email)
	return ctx.Redirect(http.StatusMovedPermanently, "/dashboard/" + user_id)
}

func dashboard_handler(ctx echo.Context) error {

	key := ctx.Param("key")
	user_data := load_user_data(key)
	t, _ := template.ParseFiles("templates/dashboard.html")
	content := make(map[string]string)
	content["email"] = user_data["email"]
	content["name"] = user_data["name"]
	content["company_name"] = user_data["company_name"]
	content["key"] = key
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.HTML(http.StatusOK, string(buffer.Bytes()))
}

func dashboard_handler_post(ctx echo.Context) error {

	// retrieving & storing form data
	key := ctx.Param("key")
	name := ctx.FormValue("name")
	company_name := ctx.FormValue("company_name")
	user_data := load_user_data(key)
	user_data["name"] = name
	user_data["company_name"] = company_name
	store_user_data(key, user_data)

	// response
	t, _ := template.ParseFiles("templates/dashboard.html")
	content := make(map[string]string)
	content["email"] = user_data["email"]
	content["name"] = user_data["name"]
	content["company_name"] = user_data["company_name"]
	content["key"] = key
	buffer := new(bytes.Buffer)
	t.Execute(buffer, content)
	return ctx.HTML(http.StatusOK, string(buffer.Bytes()))
}

func main() {

	e := echo.New()
	// ========== routing ==========
	e.GET("/", default_handler)
	e.POST("/login", login_handler)
	e.GET("/dashboard/:key", dashboard_handler)
	e.POST("/dashboard/:key", dashboard_handler_post)
	// listenning
	e.Logger.Fatal(e.Start(":8081"))
}
