package main

import (
	"github.com/DiegoUrrego4/inventory-app/cmd/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
