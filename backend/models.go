package main

type User struct {
	Username   string  `json:"username", db:"username"`
	Password   string  `json:"password", db:"password"`
	//Role       string  `json:"role"`
	//Status     string  `json:"status"`
}

/*
type Task struct {

}
*/
