package main

// import ()

var (
    sessionKey = []byte("sup3r-s3cr3t-k3y")
)

func main() {
    a := App{}
	a.Initialize(sessionKey)

	a.Run(":8080")
}
