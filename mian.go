package main

import (
	//"context"
	"database/sql"
	"fmt"
	"github.com/godror/godror"
	"os"
)

func main() {
	user := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	connectString := os.Getenv("DB_CONNECT_STRING") //for example "examplepdb_tp"
	connectionString := user + "/" + dbpassword + "@" + connectString

	fmt.Fprintf(os.Stdout, "try connecting using string : %s and driver %s ", connectionString, godror.DriverName)

	connection, err := sql.Open("godror", connectionString)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = connection.Ping()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer connection.Close()

	println("Hello, world.")
}
