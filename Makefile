.PHONY: migrate-up migrate-down test run-user run-order swagger-user swagger-order

DB_URL=postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable

# Appliquer les migrations / Применить миграции
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

# Annuler les migrations / Откатить миграции
migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down

# Lancer les tests / Запустить тесты
test:
	go test -race -cover ./...

# Lancer le user-service / Запустить user-service
run-user:
	go run cmd/user-service/main.go

# Lancer le order-service / Запустить order-service
run-order:
	go run cmd/order-service/main.go

# Générer Swagger pour user-service / Сгенерировать Swagger для user-service
swagger-user:
	swag init -g cmd/user-service/main.go -o cmd/user-service/docs

# Générer Swagger pour order-service / Сгенерировать Swagger для order-service
swagger-order:
	swag init -g cmd/order-service/main.go -o cmd/order-service/docs

# Construire avec Docker / Собрать с Docker
docker-build:
	docker-compose build

# Lancer avec Docker / Запустить с Docker
docker-up:
	docker-compose up

# Arrêter avec Docker / Остановить с Docker
docker-down:
	docker-compose down






	# Lancer tous les tests avec couverture détaillée
# Запустить все тесты с детальным покрытием
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Rapport généré dans coverage.html / Отчёт создан в coverage.html"

# Lancer les tests par module / Запустить тесты по модулям
test-user:
	go test -v -cover ./internal/user/...

test-order:
	go test -v -cover ./internal/order/...

# Créer une DB de test / Создать тестовую БД
test-db:
	docker exec -it postgres psql -U postgres -c "CREATE DATABASE postgres_test;"

# Lancer les tests avec la vraie DB / Запустить тесты с реальной БД
test-integration:
	TEST_DB_URL="postgresql://postgres:pass@localhost:5432/postgres_test?sslmode=disable" \
	go test -v -cover ./internal/.../repository/...