local-db:
	docker stop bench-spring-db || true
	docker rm bench-spring-db || true
	docker run -e POSTGRES_PASSWORD=postgres -p 1111:5432 --name bench-spring-db -d postgres:15.3
	docker start bench-spring-db

spring-setup: local-db
	docker build -f docker/spring/Dockerfile -t benchmarks-spring:latest .

spring-run: spring-setup
	docker stop spring-bench-container || true
	docker rm spring-bench-container || true
	sleep 3
	docker run -d --memory=512m --cpus=1 --network=host --name=spring-bench-container benchmarks-spring:latest

spring-bench: spring-run
	sleep 15
	ab -n 10000 -c 100 localhost:9095/api/top-entities


golang-setup:
	$(MAKE) -C go-http-server local-db
	docker build -f docker/go/Dockerfile -t benchmarks-golang:latest .

golang-run: golang-setup
	docker stop golang-bench-container || true
	docker rm golang-bench-container || true
	sleep 3
	docker run -d --memory=512m --cpus=1 --network=host --name=golang-bench-container benchmarks-golang:latest

golang-bench: golang-run
	sleep 15
	ab -n 10000 -c 100 localhost:9096/api/top-entities
