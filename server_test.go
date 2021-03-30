package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lhervi/test_postgresql2/pkg/connection"
)

func TestRunServer(t *testing.T) {

	const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"
	conn := new(connection.Postgresqlconn)
	err := conn.Connect(psqlInfo)
	t.Run("Connection test", func(t *testing.T) {
		if err != nil {
			t.Fatalf(err.Error())
		}
	})

	repo := &Postgresrepo{conn}

	handler := NewHandlers(repo)
	myapp := NewTyniServer(handler)
	t.Run("Server creationn", func(t *testing.T) {
		if myapp == nil {
			err := "the server was not created"
			t.Fatalf(err)
		}
	})

	typeMyapp := fmt.Sprint(reflect.TypeOf(myapp))
	t.Run("Server Type ok", func(t *testing.T) {
		if typeMyapp != "*main.TyniServer" {
			err := "there is a problem with the server type created"
			t.Fatalf(err)
		}
	})
}
