run:
	go mod download && docker-compose up --build

create-migration:
	# make create-migration NAME=create_some_table
	migrate create -ext sql -dir ./db/migrations -seq $(NAME)

migrate:
	migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up 1

migrate-down:
	migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable down 1

migrate-drop:
	migrate -path ./db/migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable drop
