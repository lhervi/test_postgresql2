package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgresqlconn struct {
	ConnString string
	DBConn     *sql.DB
}

//func NewPostgresqlconn() *Postgresqlconn {
func New() *Postgresqlconn {
	return &Postgresqlconn{}
}

func (pc *Postgresqlconn) Connect(psqlInfo string) error {

	if psqlInfo == "" {
		err := fmt.Errorf("the connection string is empty")
		return err
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		//err2:="Check the connection string"
		fmt.Println(err)
		return err
	}
	pc.ConnString = psqlInfo
	pc.DBConn = db
	return nil
}
