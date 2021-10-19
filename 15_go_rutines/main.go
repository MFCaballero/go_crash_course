package main

import (
	"fmt"
	"strings"
	"time"
)

func splitName(name string) {
	letters := strings.Split(name, "")

	for _, value := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(value)
	}
}

func sync(name string) {
	splitName(name)
	fmt.Println("Finished!")
}

func conc(name string) {
	go splitName(name)
	fmt.Println("Finished!")
	var wait string
	fmt.Scanln(&wait)
}

func main() {
	name := "Florencia"
	fmt.Println("Synchronous")
	sync(name)
	fmt.Println("Concurrent")
	conc(name)
}
