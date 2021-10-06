package main

import (
	"log"
	"net/http"
  //"encoding/json"
	"html/template"
)

type User struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/users", user_handler)
	port := ":8080"
	println("Server listening port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}


func handler(w http.ResponseWriter, r *http.Request) {
	//user := User{"Yerek", "Yerkinbekuly"}
  //js, _ := json.Marshal(user)

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
}

func user_handler(w http.ResponseWriter, r *http.Request){
	users := []User{User{"Bolat", "Bekbenbet"}, User{"Yerek", "Yerkinbekuly"}}
	tmpl, err := template.ParseFiles("static/users.html")
	if err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
	if err := tmpl.Execute(w, users); err != nil{
		http.Error(w, err.Error(), 400)
		return
	}
}
