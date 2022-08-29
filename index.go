package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name              string
	Age               uint16
	Money             int32
	Avg_grades, Happy float64
	Hobbies           []string
}

func (user *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is: %d and he has money: %d", user.Name, user.Age, user.Money)
}

func (user *User) setNewName(newName string) {
	user.Name = newName
}

func home(write http.ResponseWriter, request *http.Request) {
	anton := User{"Anton", 32, 100500, 98.1, 99.5, []string{"foot", "hands", "listens"}}
	// anton.setNewName("new Anton name")
	// fmt.Fprintf(write, anton.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(write, anton)
}
func contact(write http.ResponseWriter, request *http.Request) {
	anton := User{"Anton", 32, 100500, 98.1, 99.5, []string{"foot", "hands", "listens"}}
	fmt.Fprintf(write, "<h1>Contact User name"+anton.Name+"</h1>")
}

func pages() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact/", contact)
	http.ListenAndServe(":8081", nil)
}

func main() {

	//anton := User{name: "Anton", age:32, money: 100500, avg_grades: 98.1, happy: 99.5}

	pages()
}
