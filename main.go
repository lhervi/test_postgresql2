package main

import "fmt"

/*
import (
	//"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)
*/

func main() {

	myapp := new(TyniServer)
	err := myapp.Init()
	if err != nil {
		fmt.Println(err)
	}

}
