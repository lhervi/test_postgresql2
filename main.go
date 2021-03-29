package main

import (
	"fmt"

	"github.com/lhervi/test_postgresql2/pkg/connection"
)

func main() {

	const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"
	conn := new(connection.Postgresqlconn)
	err := conn.Connect(psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	repo := &Postgresrepo{conn}

	handler := NewHandlers(repo)
	myapp := NewTyniServer(handler)
	err = myapp.Init()
	if err != nil {
		fmt.Println(err)
	}
}
