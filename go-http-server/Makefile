local-db:
	docker stop bench-golang-db || true
	docker rm bench-golang-db || true
	docker run -e POSTGRES_PASSWORD=postgres -p 2222:5432 --name bench-golang-db -d postgres:15.3
	docker start bench-golang-db
