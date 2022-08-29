package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.31.132:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//установка данных

	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Anto1n', 32)")

	// if err != nil {
	// 	panic(err)
	// }

	// defer insert.Close()

	//выборка

	res, err := db.Query("SELECT `name`, `age` FROM `users`")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)

		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

	//defer res.Close()

	fmt.Println("Подключено к Базе")
}
