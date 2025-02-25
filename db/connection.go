package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error cargando el archivo .env")
    }

    dsn := "host=" + os.Getenv("DB_HOST") +
    " user=" + os.Getenv("DB_USER") +
    " password=" + os.Getenv("DB_PASSWORD") +
    " dbname=" + os.Getenv("DB_NAME") +
    " port=" + os.Getenv("DB_PORT") +
    " sslmode=" + os.Getenv("SSL_MODE")

    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Conexión exitosa a PostgreSQL")
}
