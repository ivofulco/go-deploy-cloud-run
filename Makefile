
local:
	docker compose up -d;

infra-down:
	docker compose down;

go-run:
	go run main.go

go-test:
	go test ./... -v	

local-up: local
	curl http://localhost:8080/01153000