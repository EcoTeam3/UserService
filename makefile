CURRENT_DIR=$(shell pwd)

proto-gen:
	./script/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:hamidjon4424@localhost:5432/userservice?sslmode=disable'

mig-up:
	migrate -path storage/migrations -database 'postgres://postgres:hamidjon4424@localhost:5432/userservice?sslmode=disable' -verbose up

mig-down:
	migrate -path storage/migrations -database ${DBURL} -verbose down

mig-create:
	migrate create -ext sql -dir storage/migrations -seq create_tables2

mig-insert:
	migrate create -ext sql -dir storage/migrations -seq insert_table
