package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	database, er := sql.Open("sqlserver",
		"odbc:server=localhost; user id=sa; password={Athum117343}; database = ejem")

	if er != nil {
		fmt.Println(er)
	}

	var nombre string
	var apellido string
	var id int

	database.QueryRow("select * from usuario").Scan(&id, &nombre, &apellido)

	fmt.Println("Nombre : ", nombre, "\nApellido : ", apellido)
}
