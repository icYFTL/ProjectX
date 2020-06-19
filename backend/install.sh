#!/bin/bash

go get github.com/jackc/chunkreader/
go get github.com/jackc/pgconn/
go get github.com/jackc/pgio/
go get github.com/jackc/pgpassfile
go get github.com/jackc/pgservicefile/
go get github.com/jackc/puddle/
go get github.com/gorilla/mux/

go build -o ../frontend/dist/backend_compiled -i .