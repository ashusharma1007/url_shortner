package main

import (
	"url-short/pkg"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	pkg.SetupDatabase()
	pkg.RegisterRoutes()
}
