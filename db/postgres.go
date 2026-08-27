package db

import "database/sql"
import _ "github.com/lib/pq"

func Connect(connString string) (*sql.DB, error) {
    return sql.Open("postgres", connString)
}
