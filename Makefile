docker_up:
	docker-compose up

docker_build:
	docker-compose up -d --build

docker_down:
	docker-compose down

docker_remove_vol:
	docker-compose down --volume

docker_restart:
	docker-compose restart