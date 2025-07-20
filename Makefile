db/local:
	@docker run --name local-pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@sleep 3
	@psql -f db/create_tables.sql "postgresql://postgres:postgres@localhost:5432/postgres" 
	@psql -f db/initial_seed.sql "postgresql://postgres:postgres@localhost:5432/postgres"
	#@make parse
db/local/down:
	@docker rm local-pg -f
ahab:
	@docker rm -f local-pg
react/build:
	@npm run build --prefix ./ui
	@cp -r ./ui/build ./server/
	@mv ./server/build ./server/public
serve/local:
# 	ENVIRONMENT=dev go run -mod vendor ./server/main.go ./server/password/password.go ./server/structs/structs.go ./server/cookie.go ./server/structs/sql/*.go 
	@ENVIRONMENT=dev FAKE_AUTH=true go run -mod vendor ./server/main.go
# parse:
# 	@go clean -testcache
# 	@go test -v ./server/parsing/parse_test.go -count=1
test/unit:
	@go test -v ./server/parsing ./server/cookie
# make test/upload month=9 year=2024 file=capone_test.csv
# make test/upload month=9 year=2024 capone=false file=amex_test.csv
test/upload:
	@curl -v -X POST http://localhost:8080/api/upload -F "capone=$(capone)" -F "month=$(month)" -F "year=$(year)" -F "file=@sample_files/$(year)/$(file)" -H 'Content-Type: multipart/form-data'
psql:
	@psql "postgresql://postgres:postgres@localhost:5432/postgres"
serve/ui:
	npm start --prefix ./ui
serve/node:
	node ./node/index.js
dev:
	@docker start local-pg
	@node dev.js
# make signup email=david@neurocollective.io firstName=Dude lastName=Duder password=pw
signup:
	@curl -d '{ "email": "$(email)", "lastName": "$(lastName)", "firstName": "$(firstName)", "password": "$(password)" }' -H 'Accept: application/json' -H 'Content-Type: application/json' localhost:8080/signup
fmt:
	@go fmt ./server/ ./server/constants ./server/structs ./server/password ./server/constants ./server/cookie
# parse:
# 	@go run -mod vendor ./server/parse.go
install:
	@rm -rf vendor/github.com/neurocollective/go_utils
	@cp -r ../go_utils vendor/github.com/neurocollective/
	@rm -rf vendor/github.com/neurocollective/go_utils/.git
db/dump:
	pg_dump -f ./dumps/dump.sql -d postgres -h localhost -p 5432 -U postgres
db/restore:
	psql -f ./dumps/dump.sql "postgresql://postgres:postgres@localhost:5432/postgres"
db/rebuild:
	@docker run --name local-pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
	@sleep 3
	make db/restore
exec/psql:
	docker exec -it local-pg bash "/db_scripts/setup.sh"
test:
	@echo "hi there $$(date)"
make db/dump/upload:
	aws s3 cp ./dumps/dump.sql s3://neurocollective/dumps/dump.sql $$ENDPOINT
	aws s3 cp ./dumps/dump.sql s3://neurocollective/dumps/dump.sql $$VENDPOINT
