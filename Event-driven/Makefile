
start:
	docker-compose up -d
	cd kong && make kong-postgres

startwb:
	docker-compose up -d --build
	cd kong && make kong-postgres

clean:
	docker-compose kill
	docker-compose rm -f