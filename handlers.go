package main

import (
	"fmt"
	"net/http"

	"github.com/lhervi/test_postgresql2/pkg/connection"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func GetAll(c *gin.Context) {

	db := new(connection.Postgresqlconn)

	err := db.Connect(psqlInfo)

	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		fmt.Print(err.Error())
		return
	}

	var (
		user  User
		users []User
	)

	rows, err := db.DBConn.Query("SELECT id, name, lastname, email, role FROM UserInfo")
	if err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		fmt.Print(err.Error())
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Role)
		users = append(users, user)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNoContent)
			fmt.Print(err.Error())
			return
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}
