package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect_InvalidURL(t *testing.T) {
	// sql.Open ne se connecte pas réellement, juste valide le format
	// sql.Open не подключается реально, только валидирует формат
	db, err := Connect("postgresql://invalid:invalid@localhost:9999/none?sslmode=disable")
	assert.NoError(t, err)
	assert.NotNil(t, db)
	defer db.Close()
}

func TestConnect_RealDB(t *testing.T) {
	db, err := Connect("postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		t.Skip("DB non disponible / БД недоступна")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		t.Skip("DB non joignable / БД недоступна")
	}
	assert.NoError(t, err)
}
