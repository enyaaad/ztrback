package main

import (
	server "austem"
	"austem/migrations"
)

func main() {
	migrations.Migrator()
	server.StartAPI()
}
