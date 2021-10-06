package main

import(
  "fmt"
  "net/http"
  "html/template"
)

type User struct {
  Name string
  Age uint16
  Money int16
  Status string
  Hobbies []string
}

type Contact struct{
  Name string
  Number string
  Email string
}

// func (u *User) getAllInfo()string ссылка қоюға болады
func (u User) getAllInfo()string{
  return fmt.Sprintf("User name is: %s. He is %d and he has money"+"equal: %d", u.Name, u.Age, u.Money)
}
//ссылка міндетті түрде қою керек
func (u *User) setNewName(newName string){
  u.Name = newName
}

func handleRequest(){
  fmt.Println("Listening server ...")
  http.HandleFunc("/", homePage)
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request){
  yerek := User{"Yerek", 21, 0, "programmer", []string{"Running", "Cycling", "Coding"}}
  tmpl, _ := template.ParseFiles("templates/homepage.html")
  tmpl.Execute(w, yerek)
}

func contacts_page(w http.ResponseWriter, r *http.Request){
  contacts := Contact{"Yerek", "+77471850499", "erkinbekly@gmail.com"}
  tmpl, _ := template.ParseFiles("templates/contact.html")
  tmpl.Execute(w, contacts)
}


func main(){
   handleRequest()
}
