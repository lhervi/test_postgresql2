package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"

func GetAll(c *gin.Context) {

	db := Postgresqlconn.NewPostgresconnn()

	//db, err := sql.Open("postgres", psqlInfo)

	var (
		user  User
		users []User
	)

	rows, err := db.Query("SELECT id, name, lastname, email, role FROM UserInfo")
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
