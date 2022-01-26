package main

import (
	"sources/data"
	"sources/server"
)

func main() {
	data.InitDB()
	server.Run()
}
