up:
	docker compose up -d
down:
	docker compose down
re-exec:
	docker exec -it front bash
module-install:
	docker compose -f docker-compose.setup.yml run --rm module-install
react-install:
	docker compose -f docker-compose.setup.yml run --rm installer
react-install-dev:
	docker compose -f docker-compose.setup.yml run --rm installer-dev
node-module-clean:
	docker volume rm react-ptjcist_node-modules-volume