// cmd/smartcrypto/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartcrypto/internal/smartcrypto"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartcrypto.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
