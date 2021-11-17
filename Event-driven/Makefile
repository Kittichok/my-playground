
start:
	docker-compose up -d
	cd product && docker-compose -f docker-compose.hotreload.yml up -d
	cd booking && docker-compose -f docker-compose.hotreload.yml up -d
	cd payment && docker-compose -f docker-compose.hotreload.yml up -d
	cd kong && make kong-postgres

# startwb:
# 	cd pproduct && docker-compose -f docker-compose.hotreload.yml up -d --build
# 	cd kong && make kong-postgres
down:
	docker-compose down
	cd product && docker-compose -f docker-compose.hotreload.yml down
	cd booking && docker-compose -f docker-compose.hotreload.yml down
	cd payment && docker-compose -f docker-compose.hotreload.yml down
	cd kong && make clean

# test:
# 	echo "not implemented yet"

clean:
	docker-compose kill
	docker-compose rm -f

kong sync:
	cd kong/config && deck sync

    # docker run -it --rm \
    # --network backend \
    # -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
    # bitnami/kafka:latest kafka-topics.sh --list  --bootstrap-server kafka:9092

# kafka_test:
#     docker run -it --rm \ 
# 		--network backend \ 
# 		-e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
#     bitnami/kafka:latest kafka-console-producer.sh \
# 		--topic booking --broker-list kafka:9092

# docker run -it --rm --network backend -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 \
# bitnami/kafka:latest kafka-console-consumer.sh \
# --topic booking --from-beginning --bootstrap-server kafka:9092
    
loadtest:
	cd tests/loadtest && k6 run create-product.js
# cd tests/loadtest && k6 run booking.js
