
start:
	cd pproduct && docker-compose -f docker-compose.hotreload.yml up -d
	cd kong && make kong-postgres

# startwb:
# 	cd pproduct && docker-compose -f docker-compose.hotreload.yml up -d --build
# 	cd kong && make kong-postgres

clean:
	docker-compose kill
	docker-compose rm -f