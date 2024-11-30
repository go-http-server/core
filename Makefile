setup-service:
	docker-compose -f docker-compose.yml up -d --build
delete-service:
	docker-compose down -v
mockgen:
	mockgen -package mockgen_db -destination ./internal/database/mockgen/store.go github.com/go-http-server/core/internal/database/sqlc/ Store
