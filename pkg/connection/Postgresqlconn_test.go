package connection

import (
	"testing"
)

const psqlInfo string = "postgresql://dbuser:dbuserpassword@localhost:5432/users"
const psqlErr string = "postgresql://dbuser:dbuserpassword@localhost:5432/noDb"

func TestPostgresqlconn(t *testing.T) {
	var tests = []struct {
		testName string
		conn     string
		result   string
	}{
		{"Connection string ok", psqlInfo, ""},
		{"No connection string", "", "the connection string is empty"},
		{"Wrong connection string", psqlErr, "check the connnection string"},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			conn := New()
			err := conn.Connect(test.conn)
			if err != nil {
				comp := err.Error()
				if comp != test.result {
					t.Errorf("the error: %v", err)
				}
			}

		})
	}
}
