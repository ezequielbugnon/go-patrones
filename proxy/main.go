package main

import (
	"fmt"
	"time"

	"github.com/ezequielbugnon/go-patrones/proxy/local"
	"github.com/ezequielbugnon/go-patrones/proxy/remote"
)

var loc local.Proxy

func main() {
	//proxy example

	loc = local.New(remote.New("http://something", 8080, "usuario", "123456"))

	GetById(1)
	GetById(0)
	GetById(1)
	GetById(0)
	GetAll()
}

func GetById(ID uint) {
	start := time.Now()
	fmt.Printf("%+v", loc.GetByID(ID))
	elapsed := time.Since(start)
	fmt.Printf("tiempo usado: %v\n", elapsed)
}

func GetAll() {
	start := time.Now()
	fmt.Printf("%+v", loc.GetAll())
	elapsed := time.Since(start)
	fmt.Printf("tiempo usado: %v\n", elapsed)
}
