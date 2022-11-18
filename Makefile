docker_up:
	docker-compose up -d

docker_down:
	docker-compose down

docker_remove_vol:
	docker-compose down --volume

docker_restart:
	docker-compose restart