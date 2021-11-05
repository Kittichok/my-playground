kong-postgres:
	COMPOSE_PROFILES=database KONG_DATABASE=postgres docker-compose up -d
	sleep 10
	cd config && deck sync

kong-dbless:
	docker-compose up -d

sync:
	cd config && deck sync

clean:
	docker-compose kill
	docker-compose rm -f