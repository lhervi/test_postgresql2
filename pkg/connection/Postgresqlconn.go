package connection

import (
	"database/sql"
	//"fmt"
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

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	pc.ConnString = psqlInfo
	pc.DBConn = db
	return nil
}
