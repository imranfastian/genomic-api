package main

import (
	"genomic-api/config"
	"genomic-api/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
