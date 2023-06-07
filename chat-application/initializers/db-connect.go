package initializers

import (
	users "chat-application/sqlc-models"
	"chat-application/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var Db *users.Queries

func ConnectToDb() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	conn, connectErr := sql.Open("postgres", dsn)
	utils.CheckError(connectErr, "Error connecting to database")
	fmt.Println("Successfully connected to database")
	db := users.New(conn)
	Db = db
}
