package main

import "sources/data"

func main() {
	data.InitDB()
	data.StartServer()
}
