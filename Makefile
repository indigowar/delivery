
gen:
	go generate ./...

test: gen
	set -o pipefail && go test ./... -json  | tparse -all
	# go test ./...

check: test
	go vet ./...
	staticcheck ./...

up: gen
	docker-compose -f build/local-compose.yaml --env-file .env up --build --remove-orphans
