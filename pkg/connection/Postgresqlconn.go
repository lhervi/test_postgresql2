package connection

import (
	"database/sql"
	//"fmt"
)

type Postgresqlconn struct {
	connString string
	db         *sql.DB
}

func NewPostgresqlconn() *Postgresqlconn {
	return &Postgresqlconn{}
}

func (pc *Postgresqlconn) Connect(psqlInfo string) error {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	pc.connString = psqlInfo
	pc.db = db
	return nil
}
