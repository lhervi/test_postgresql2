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
		lenErr   bool
	}{
		{"Connection string ok", psqlInfo, true},
		{"No connection string", "", false},
		{"Wrong connection string", psqlErr, false},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			conn := New()
			err := conn.Connect(test.conn)
			if err != nil {
				t.Errorf("the error: %v", err)
			}
		})
	}
}
