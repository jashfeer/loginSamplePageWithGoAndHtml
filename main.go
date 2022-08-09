package main

import (
	"login/controllers"
	"net/http"

)

func main() {

	http.HandleFunc("/account", controllers.Index)
	http.HandleFunc("/account/index", controllers.Index)
	http.HandleFunc("/account/login", controllers.Login)
	http.HandleFunc("/account/welcome", controllers.Welcome)
	http.HandleFunc("/account/logout", controllers.Logout)

	http.ListenAndServe(":3000",nil)

}
