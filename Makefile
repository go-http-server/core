setup-service:
	docker-compose -f docker-compose.yml up -d --build
delete-service:
	docker-compose down -v
mockgen:
	mockgen -package mockgen_db -destination ./internal/database/mockgen/store.go github.com/go-http-server/core/internal/database/sqlc/ Store
gen-docs: 
	npx @redocly/cli build-docs ./docs/swagger.yaml --output=./docs/api-docs.html
gen-via-openapi:
	swag init --dir ./api --parseVendor=true --output=./docs --generalInfo server.go  --pdl=3 --dir ./cmd/
