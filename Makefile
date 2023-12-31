db/local:
	@docker run --name local-pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@sleep 2
	@psql -f db/create_tables.sql "postgresql://postgres:postgres@localhost:5432/postgres" 
	@psql -f db/initial_seed.sql "postgresql://postgres:postgres@localhost:5432/postgres"
ahab:
	@docker rm -f local-pg
test/parse:
	go test -v ./src/parsing
react/build:
	@npm run build --prefix ./ui
	@cp -r ./ui/build ./src/
	@mv ./src/build ./src/public
server/local:
	go run ./src/main.go
parse/test:
	@go test -v ./src/parsing
serve/local:
	go run src/main.go
psql:
	@psql "postgresql://postgres:postgres@localhost:5432/postgres"
