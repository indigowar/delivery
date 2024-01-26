
gen:
	go generate ./...

test: gen
	go test ./...

up: gen
	docker-compose -f build/local-compose.yaml --env-file .env up --build --remove-orphans
