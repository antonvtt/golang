package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Id        uint16
	Title     string
	Anons     string
	Full_text string
}

var articles = []Article{}
var post = Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)

}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)

}

func save_article(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(192.168.31.132:3306)/golang")
		if err != nil {
			panic(err)
		}

		//установка данных

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s','%s', '%s')", title, anons, full_text))

		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/posts/", 301)
	}
}

func allposts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/posts.html", "templates/header.html", "templates/footer.html")
	db, err := sql.Open("mysql", "root:root@tcp(192.168.31.132:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles` ORDER BY `id` DESC")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)

		if err != nil {
			panic(err)
		}

		articles = append(articles, post)

		//fmt.Println(fmt.Sprintf("Post: %s with age %s", post.Id, post.Title))
	}
	fmt.Println("Подключено к Базе")
	t.ExecuteTemplate(w, "allposts", articles)
	//defer res.Close()

}

func thispost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "ID: %v\n", vars["id"])

	t, err := template.ParseFiles("templates/post.html", "templates/header.html", "templates/footer.html")
	db, err := sql.Open("mysql", "root:root@tcp(192.168.31.132:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s' ", vars["id"]))

	if err != nil {
		panic(err)
	}
	post = Article{}
	for res.Next() {
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)

		if err != nil {
			panic(err)
		}

		post = post
	}

	//fmt.Println(fmt.Sprintf("Post: %s with age %s", post.Id, post.Title))

	fmt.Println("Подключено к Базе")
	t.ExecuteTemplate(w, "post", post)
	//defer res.Close()

}

func handleFunction() {

	router := mux.NewRouter()

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/create/", create).Methods("GET")
	router.HandleFunc("/posts/", allposts).Methods("GET")
	router.HandleFunc("/post/{id:[0-9]+}", thispost).Methods("GET")
	router.HandleFunc("/save_article/", save_article).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe("golangtest.local:80", nil)
}

func main() {
	handleFunction()
}
