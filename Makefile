local-up:
	docker compose up -d;

infra-down:
	docker compose down;

docker-clean-up:
	docker rm $(docker ps -aq) -f	