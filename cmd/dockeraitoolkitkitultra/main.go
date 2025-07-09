// cmd/dockeraitoolkitkitultra/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"dockeraitoolkitkitultra/internal/dockeraitoolkitkitultra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := dockeraitoolkitkitultra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
