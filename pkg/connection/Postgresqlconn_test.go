package connection

const pgaddr string = "postgresql://user:password123@localhost:5432/shore_test"

/*
func TestConnect(T *testing.T) {

	pgrep := New()


	ctx := context.Background()

	//pgCloser, err := pgrep.Connect(ctx, pgaddr)
	pgCloser, err :=

	if err != nil {
		T.Fatalf(err.Error())
	}
	defer pgCloser()

}
*/
