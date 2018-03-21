package main

import (
	"github.com/Vexth/GolangDemo/new"
	"github.com/Vexth/GolangDemo/router"
)

func main() {
	new.HandlerFunc("/hello", router.SayHello)
	new.HandlerFunc("/bye", router.SayBye)

	new.StartServer(":8080", "/staic/")
}
