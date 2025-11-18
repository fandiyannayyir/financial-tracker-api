package config

import (
	"database/sql"
	"financial-tracker-api/utils"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := utils.GetEnv("DB_HOST", "localhost")
	portStr := utils.GetEnv("DB_PORT", "5432")
	user := utils.GetEnv("DB_USER", "postgres")
	password := utils.GetEnv("DB_PASSWORD", "")
	database := utils.GetEnv("DB_NAME", "financial_tracker")
	sslmode := utils.GetEnv("DB_SSL_MODE", "disable")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid port number:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, database, sslmode)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully! 🚀")
}
