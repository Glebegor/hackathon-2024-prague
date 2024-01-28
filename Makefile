rundb:
	sudo docker run --name=ton-event-db -e POSTGRES_PASSWORD="123321" -p 5555:5432 --rm -d postgres
migrate-up:
	migrate -path ./backend/schema -database "postgres://postgres:123321@localhost:5555/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./backend/schema -database "postgres://postgres:123321@localhost:5555/postgres?sslmode=disable" down