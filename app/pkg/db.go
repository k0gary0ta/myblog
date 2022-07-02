package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
)

func GenerateULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	// fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func OpenDB() *sql.DB {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("error opening env file: ", envErr)
	}

	cfg := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               os.Getenv("MYSQL_DATABASE"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("db.Ping failed:", err)
	}

	fmt.Println("db connected")
	return db
	// defer db.Close()
}
