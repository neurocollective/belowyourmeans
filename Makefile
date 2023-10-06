db/local:
	@docker run --name local-pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@psql "postgresql://postgres:postgres@localhost:5432/postgres" -f db/create_tables.sql
	@psql "postgresql://postgres:postgres@localhost:5432/postgres" -f db/initial_seed.sql
ui/build:
	@npm run build --prefix ./ui
	@cp -r ./ui/build ./src/
	@mv ./src/build ./src/public
server/local:
	go run ./src/main.go 