BENCH_REQUESTS ?= 1000

local-db:
	docker stop bench-golang-db || true
	docker rm bench-golang-db || true
	docker run -e POSTGRES_PASSWORD=postgres -p 2222:5432 --name bench-golang-db -d postgres:15.3
	docker start bench-golang-db

setup: local-db
	docker build -f $(CURDIR)/../docker/go/Dockerfile -t benchmarks-golang:latest $(CURDIR)/..

run: setup
	docker stop golang-bench-container || true
	docker rm golang-bench-container || true
	sleep 3
	docker run -d --memory=512m --cpus=1 --network=host --name=golang-bench-container benchmarks-golang:latest

bench: run
	sleep 15
	ab -n $(BENCH_REQUESTS) -c 100 localhost:9096/api/top-entities

all: bench