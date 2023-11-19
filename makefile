null:
	@:

run:
	go run ./app

migration:
	go run ./tools/migrations


docker-up:
	docker compose up --detach