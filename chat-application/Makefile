postgres-init:
	sudo docker run --name postgres15 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Simform@123 -p 5432:5432 -d postgres:15-alpine

postgres:
	sudo docker exec -it postgres15 psql -U postgres

create-db:
	sudo docker exec -it postgres15 createdb --username=postgres --owner=postgres chat-app

drop-db:
	sudo docker exec -it postgres15 dropdb chat-app

migrate-up:
	 migrate -database ${POSTGRESQL_URL} -path db/migrations -verbose up

migrate-down:
	 migrate -database ${POSTGRESQL_URL} -path db/migrations -verbose down

migration-fix:
	migrate -database ${POSTGRESQL_URL} -path db/migrations -verbose force <VERSION>
sqlc:
	sqlc generate
.PHONY: postgres-init postgres create-db drop-db