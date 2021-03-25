package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"

func GetAll(c *gin.Context) {

	db := new(Postgresqlconn)
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

	rows, err := db.dbConn.Query("SELECT id, name, lastname, email, role FROM UserInfo")
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

func main() {

	router := gin.Default()

	router.GET("/all", GetAll) // **************  Get all users  **************

	router.Run(":3020")
}
