package main

import (
	"pdb/database"
	"pdb/router"
)

func main() {
	database.SetupDB()
	database.StartScheduler()

	r := router.Initialize()

	_ = r.Run()
}
