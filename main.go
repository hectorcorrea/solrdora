package main

import (
	"gosiah/web"
)

func main() {
	address := "localhost:9001"
	web.StartWebServer(address)
}
