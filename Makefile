docker_up:
	docker-compose up -d

docker_down:
	docker-compose down

docker_remove_vol:
	docker-compose down --volume

docker_restart:
	docker-compose restart

db_seed:
	docker-compose exec db mysql -u root < db/schema.sql