setup-service:
	docker-compose -f docker-compose.yml up -d --build
delete-service:
	docker-compose down -v
