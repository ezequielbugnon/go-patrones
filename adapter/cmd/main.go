package main

import (
	"fmt"
	"log"

	"github.com/ezequielbugnon/go-patrones/adapter/adapter"
)

func main() {
	var t string
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatalf("no se pudo leer el medio de transporte: %v", err)
	}
	transporte := adapter.Factory(t)
	transporte.Mover()
}
