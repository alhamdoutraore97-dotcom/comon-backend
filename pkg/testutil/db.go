package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// SetupTestDB crée une connexion de test à PostgreSQL
// SetupTestDB создаёт тестовое подключение к PostgreSQL
func SetupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// Récupère l'URL depuis les variables d'environnement
	// Получает URL из переменных окружения
	connStr := os.Getenv("TEST_DB_URL")
	if connStr == "" {
		connStr = "postgresql://postgres:1234@localhost:5432/postgres_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Impossible d'ouvrir la DB / Не удалось открыть БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Skipf("PostgreSQL non disponible, test ignoré / PostgreSQL недоступен, тест пропущен: %v", err)
	}

	// Nettoie les tables avant chaque test
	// Очищает таблицы перед каждым тестом
	cleanup := func() {
		db.Exec("TRUNCATE TABLE orders RESTART IDENTITY CASCADE")
		db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	}
	cleanup()
	t.Cleanup(cleanup)

	return db
}

// CreateTestUser insère un utilisateur de test
// CreateTestUser вставляет тестового пользователя
func CreateTestUser(t *testing.T, db *sql.DB, name, email string) int {
	t.Helper()
	var id int
	err := db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		name, email,
	).Scan(&id)
	if err != nil {
		t.Fatalf("Impossible de créer l'utilisateur / Не удалось создать пользователя: %v", err)
	}
	return id
}

// RunMigrations applique les migrations minimales pour les tests
// RunMigrations применяет минимальные миграции для тестов
func RunMigrations(t *testing.T, db *sql.DB) {
	t.Helper()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            created_at TIMESTAMP DEFAULT NOW()
        )`,
		`CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
            product TEXT NOT NULL,
            quantity INT NOT NULL CHECK (quantity > 0),
            status TEXT DEFAULT 'pending',
            created_at TIMESTAMP DEFAULT NOW()
        )`,
	}
	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			t.Fatalf("Migration échouée / Миграция не удалась: %v\nQuery: %s", err, q)
		}
	}
	fmt.Println("✅ Migrations de test appliquées / Тестовые миграции применены")
}
