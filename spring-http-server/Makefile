BENCH_REQUESTS ?= 1000

local-db:
	docker stop bench-spring-db || true
	docker rm bench-spring-db || true
	docker run -e POSTGRES_PASSWORD=postgres -p 1111:5432 --name bench-spring-db -d postgres:15.3
	docker start bench-spring-db

setup: local-db
	docker build -f $(CURDIR)/../docker/spring/Dockerfile -t benchmarks-spring:latest $(CURDIR)/..

run: setup
	docker stop spring-bench-container || true
	docker rm spring-bench-container || true
	sleep 3
	docker run -d --memory=512m --cpus=1 --network=host --name=spring-bench-container benchmarks-spring:latest

bench: run
	sleep 15
	ab -n $(BENCH_REQUESTS) -c 100 localhost:9095/api/top-entities

all: bench