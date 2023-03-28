package main

import (
	"context"

	"github.com/kuadran-code/go-simple-rest/database"
	"github.com/kuadran-code/go-simple-rest/internal/server"
)

func main() {
	db := database.NewSqlDB(
		"postgres",    //driver
		"localhost",   //host
		"5433",        //port
		"postgres",    //username
		"password",    //password
		"simple_rest", //database
	).ORM()

	ht := server.NewServer(db)
	defer ht.Done()
	ht.Run(context.Background(), 5000)
}
