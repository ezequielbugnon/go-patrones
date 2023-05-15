package main

import (
	"fmt"
	"sync"

	"github.com/ezequielbugnon/go-patrones/singleton/client"
	"github.com/ezequielbugnon/go-patrones/singleton/singleton"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client.IncrementAgeTwo()
		}()
		go func() {
			defer wg.Done()
			client.IncrementAge()
		}()
	}

	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
