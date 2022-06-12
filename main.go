package main

import (
	"log"

	"github.com/takadev15/onlineshop-api/config"
	"github.com/takadev15/onlineshop-api/router"
)

func main() {
  config.DBInit()
  r := router.RoutesList()
  log.Fatal(r.Run(":3030"))
}
