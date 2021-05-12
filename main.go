package main

import (
	"pdb/database"
	"pdb/router"
)

func main() {
	database.SetupDB()

	r := router.Initialize()

	_ = r.Run()
}
