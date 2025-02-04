local-db:
	docker stop bench-spring-db || true
	docker rm bench-spring-db || true
	docker run -e POSTGRES_PASSWORD=postgres -p 1111:5432 --name bench-spring-db -d postgres:15.3
	docker start bench-spring-db
