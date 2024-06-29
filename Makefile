up:
	docker compose up -d
down:
	docker compose down
re-exec:
	docker exec -it front bash
go-exec:
	docker exec -it backend bash
sql-exec:
	docker exec -it mysql bash
module-install:
	docker compose -f docker-compose.setup.yml run --rm module-install
react-install:
	docker compose -f docker-compose.setup.yml run --rm installer
react-install-dev:
	docker compose -f docker-compose.setup.yml run --rm installer-dev
node-module-clean:
	docker volume rm react-ptjcist_node-modules-volume
clear-db:
	docker volume rm react-ptjcist_cist_ptj_data