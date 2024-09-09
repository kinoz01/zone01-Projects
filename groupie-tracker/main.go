package main

import (
	apiserver "groupie/API"
	"groupie/server"
)

func main() {
	go server.Serve()
	go apiserver.Serve()

	select{}
}
