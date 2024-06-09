db/local:
	@docker run --name local-pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@sleep 2
	@psql -f db/create_tables.sql "postgresql://postgres:postgres@localhost:5432/postgres" 
	@psql -f db/initial_seed.sql "postgresql://postgres:postgres@localhost:5432/postgres"
db/local/down:
	@docker rm local-pg -f
ahab:
	@docker rm -f local-pg
test/parse:
	go test -v ./src/parsing
react/build:
	@npm run build --prefix ./ui
	@cp -r ./ui/build ./src/
	@mv ./src/build ./src/public
serve/local:
# 	ENVIRONMENT=dev go run -mod vendor ./src/main.go ./src/password/password.go ./src/structs/structs.go ./src/cookie.go ./src/structs/sql/*.go 
	ENVIRONMENT=dev go run -mod vendor ./src/main.go
parse/test:
	@go test -v ./src/parsing
psql:
	@psql "postgresql://postgres:postgres@localhost:5432/postgres"
serve/ui:
	npm start --prefix ./ui
serve/node:
	node ./node/index.js
dev:
	@docker start local-pg
	@node dev.js
signup:
	@curl -d '{ "email": "$(email)", "lastName": "$(lastName)", "firstName": "$(firstName)", "password": "$(password)" }' -H 'Accept: application/json' -H 'Content-Type: application/json' localhost:8080/signup
fmt:
	go fmt ./src
parse:
	go run -mod vendor ./src/parse.go
install:
	rm -rf vendor/github.com/neurocollective/go_utils
	cp -r ../go_utils vendor/github.com/neurocollective/
	rm -rf vendor/github.com/neurocollective/go_utils/.git
