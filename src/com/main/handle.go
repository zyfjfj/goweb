package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHellowName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "你好！看见了吗？")
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../html/login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
