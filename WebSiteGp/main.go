package main

import (
  "fmt"
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/gorilla/mux"
)


type Article struct{
  Id uint16
  Title, Anons, FullText string
}

var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/index.html","templates/header.html","templates/footer.html")
  if err != nil{
    fmt.Fprintf(w, err.Error())
  }
  connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
      panic(err)
  }
  defer db.Close()

  //Таңдау
  res, err := db.Query("SELECT * FROM Post")
  if err != nil{
    panic(err)
  }
  posts = []Article{}
  for res.Next(){
    var post Article
    err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)
    if err != nil{
      panic(err)
    }

    posts = append(posts, post)
    //fmt.Println(fmt.Sprintf("Post: %s with id: %d", post.Title, post.Id))
  }


  t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request){
  t, err := template.ParseFiles("templates/create.html","templates/header.html","templates/footer.html")

  if err != nil{
    fmt.Fprintf(w, err.Error())
  }
  t.ExecuteTemplate(w, "create", nil)
}

func save_article(w http.ResponseWriter, r *http.Request){
  title := r.FormValue("title")
  anons := r.FormValue("anons")
  full_text := r.FormValue("full_text")

  if title =="" || anons == "" || full_text == ""{
    fmt.Fprintf(w, "Не все данные заполнены")
  }else{
    connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    //Установка данных
    q := `insert into Post(title, anons, full_text) values($1, $2, $3)`
    insert, err := db.Query(q, title, anons, full_text)
    if err != nil {
        panic(err)
    }
    defer insert.Close()
    http.Redirect(w, r, "/", 301)
  }


}

func show_post(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  t, err := template.ParseFiles("templates/show.html","templates/header.html","templates/footer.html")
  if err != nil{
    fmt.Fprintf(w, err.Error())
  }
  //Мәліметтер қорына қосылу
  connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
      panic(err)
  }
  defer db.Close()

  res, err := db.Query(fmt.Sprintf("SELECT * FROM Post WHERE id = '%s'", vars["id"]))
  if err != nil{
    panic(err)
  }
  showPost = Article{}
  for res.Next(){
    var show_post Article
    err = res.Scan(&show_post.Id, &show_post.Title, &show_post.Anons, &show_post.FullText)
    if err != nil{
      panic(err)
    }
    showPost = show_post
  }

  t.ExecuteTemplate(w, "show", showPost)


}

func handleFunc(){
  rtr := mux.NewRouter()
  rtr.HandleFunc("/", index).Methods("GET")
  rtr.HandleFunc("/create", create).Methods("GET")
  rtr.HandleFunc("/save_article", save_article).Methods("POST")
  rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

  http.Handle("/", rtr)

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.ListenAndServe(":4040", nil)
}

func main(){
  handleFunc()
}
