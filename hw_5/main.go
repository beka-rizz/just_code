package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	_, err := connect()
	if err != nil {
		return
	}
	log.Println("successfully connected to db")

	r := gin.Default()

	r.GET("/hi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi"})
	})

	log.Fatal(r.Run())
}

func connect() (*sql.DB, error) {
	connString := "host=hw5_db user=postgres password=postgres123 dbname=hw_5 sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("connection error: ", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping error: ", err.Error())
		return nil, err
	}
	return db, nil
}
