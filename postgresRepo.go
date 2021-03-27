package main

import (
	"github.com/lhervi/test_postgresql2/pkg/connection"
)

const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"

type Postgresrepo struct {
	pgconn *connection.Postgresqlconn
}

func NewPostgresrepo(pgcon *connection.Postgresqlconn) *Postgresrepo {
	return &Postgresrepo{
		pgconn: pgcon,
	}
}

func (pr *Postgresrepo) GetAll() ([]User, error) {

	db := new(connection.Postgresqlconn)
	err := db.Connect(psqlInfo)
	if err != nil {
		return nil, err
	}

	var (
		user  User
		users []User
	)

	q := `SELECT id, name, lastname, email, role FROM UserInfo`

	rows, err := db.DBConn.Query(q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Email, &user.Role)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}
	return users, err
}
