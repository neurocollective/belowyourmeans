db/local:
	@docker run --name local-pg -e POSTGRES_PASSWORD=postgres -d postgres
ui/build:
	@npm run build --prefix ./ui
	@cp -r ./ui/build ./src/
	@mv ./src/build ./src/public
