package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	sessionKey = []byte("sup3r-s3cr3t-k3y")
	postgres   = "user=projectx password=passpass dbname=projectx"
)

func main() {
	_app := App{}
	_app.Initialize(sessionKey, postgres)
	_port := "8000"

	///ARGS
	if len(os.Args) > 1 {
		if strings.Contains(os.Args[2], "port=") {
			_port = strings.Split(os.Args[2], "=")[1]
		}
		if strings.Contains(os.Args[2], "-p") {
			if _, err := strconv.Atoi(os.Args[3]); err == nil {
				_port = os.Args[3]
			}
		}
	}

	_app.Run(":" + _port)
}
