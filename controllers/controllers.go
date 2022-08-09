package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var user string
var store = sessions.NewCookieStore([]byte(user))

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("viewPage/index.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	if username == "jashfeer" || username == "abcd" && password == "123" {
		user = username
		session, _ := store.Get(request, user)
		session.Values["username"] = username
		session.Save(request, response)
		http.Redirect(response, request, "/account/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"err": "Invalid username or password",
		}
		tmp, _ := template.ParseFiles("viewPage/index.html")
		tmp.Execute(response, data)
	}
}

func Welcome(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, user)
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	if username != nil {
		tmp, _ := template.ParseFiles("viewPage/welcome.html")
		tmp.Execute(response, data)
	} else {
		http.Redirect(response, request, "/account/index", http.StatusSeeOther)
	}

}
func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, user)
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/account/index", http.StatusSeeOther)
}
