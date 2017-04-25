package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, er := sql.Open("mysql",
		"root:athum117343@tcp(127.0.0.1:3306)/red_social")
	if er != nil {
		log.Fatal(er)
	}
	defer database.Close()

	var nom string
	var email string
	var ape string
	var nou string
	var id int

	sql := "select * from users"

	err := database.QueryRow(sql).Scan(&id, &nom, &nou, &ape, &email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nom, nou, ape)

}
