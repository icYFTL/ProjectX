package main

// import ()

var (
    sessionKey = []byte("sup3r-s3cr3t-k3y")
    postgres = "user=projectx password=passpass dbname=projectx"
)

func main() {
    a := App{}
	a.Initialize(sessionKey, postgres)

	a.Run(":8080")
}
