package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kevinhartarto/workframe/internal/database"
)

var ctx = context.Background()

func main() {
	db 	:= database.StartDB()
	app := 

	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("server listening on port:%v", addr)

	log.Fatal(app.Listen(addr))
	db.Close()
}
