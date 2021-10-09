package main

import (
	"context"
	"fmt"
	"time"
	"os"

	"insta_backend/database"
	"insta_backend/requestHandler"
)


func main(){
	fmt.Println("running")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client := database.Database(ctx)
	
	var dbName string = os.Getenv("MONGO_DBNAME")
	db := client.Database(dbName) 
	
	defer cancel()
	defer client.Disconnect(ctx)
	requestHandler.RequestHandler(db, ctx)
}