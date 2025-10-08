package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"genomic-api/config"
	"genomic-api/routes"
)

// init logging + metrics registration
func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}) // human-readable in dev
}

func main() {
	config.InitDB()
	defer config.CloseDB()

	r := routes.SetupRouter()

	fmt.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
