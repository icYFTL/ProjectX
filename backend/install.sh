#!/bin/bash

#go get https://github.com/jackc/chunkreader/
#go get https://github.com/jackc/pgconn/
#go get https://github.com/jackc/pgio/
#go get https://github.com/jackc/pgpassfile
#go get https://github.com/jackc/pgservicefile/
#go get https://github.com/jackc/puddle/
#go get https://github.com/gorilla/mux/

go get .
go build -o ../frontend/dist/backend_compiled -i .