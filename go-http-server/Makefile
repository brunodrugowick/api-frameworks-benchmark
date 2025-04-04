BENCH_REQUESTS ?= 1000
CONTAINER_MEM ?= 512m
CONTAINER_CPUS ?= 2

local-db: clean-db
	docker run -e POSTGRES_PASSWORD=postgres -p 2222:5432 --name bench-golang-db -d postgres:15.3
	docker start bench-golang-db

setup: local-db
	docker build -f $(CURDIR)/../docker/go/Dockerfile -t benchmarks-golang:latest $(CURDIR)/..

run: setup clean-container
	docker run -d --memory=$(CONTAINER_MEM) --cpus=$(CONTAINER_CPUS) --network=host --name=golang-bench-container benchmarks-golang:latest

bench: run
	sleep 15
	ab -n $(BENCH_REQUESTS) -c 100 localhost:9096/api/top-entities

clean: clean-db clean-container

clean-db:
	docker stop bench-golang-db || true
	docker rm bench-golang-db || true
	sleep 2

clean-container:
	docker stop golang-bench-container || true
	docker rm golang-bench-container || true
	sleep 2

all: bench clean
