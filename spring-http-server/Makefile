BENCH_REQUESTS ?= 1000
CONTAINER_MEM ?= 512m
CONTAINER_CPUS ?= 2

local-db: clean-db
	docker run -e POSTGRES_PASSWORD=postgres -p 1111:5432 --name bench-spring-db -d postgres:15.3
	docker start bench-spring-db

setup: local-db
	docker build -f $(CURDIR)/../docker/spring/Dockerfile -t benchmarks-spring:latest $(CURDIR)/..

run: setup clean-container
	docker run -d --memory=$(CONTAINER_MEM) --cpus=$(CONTAINER_CPUS) --network=host --name=spring-bench-container benchmarks-spring:latest

bench: run
	sleep 15
	ab -n $(BENCH_REQUESTS) -c 100 localhost:9095/api/top-entities

clean: clean-db clean-container

clean-db:
	docker stop bench-spring-db || true
	docker rm bench-spring-db || true
	sleep 2

clean-container:
	docker stop spring-bench-container || true
	docker rm spring-bench-container || true
	sleep 2

all: bench clean
