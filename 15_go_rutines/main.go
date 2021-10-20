package main

import (
	"fmt"
	"strings"
	"time"
)

func sync(name string) {
	letters := strings.Split(name, "")
	fmt.Println("Series 1")
	for _, value := range letters {
		time.Sleep(1 * time.Second)
		fmt.Println(value)
	}
	fmt.Println("Series 2")
}

func conc(name string) {
	letters := strings.Split(name, "")
	fmt.Println("Series 1")
	go func() {
		for _, value := range letters {
			fmt.Println(value)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Done go routine")
	}()
	fmt.Println("Series 2")
	time.Sleep(1 * time.Second * time.Duration(len(letters)))
}

func channel(name string) {
	letters := strings.Split(name, "")
	done := make(chan bool)
	fmt.Println("Series 1")
	go func() {
		for _, value := range letters {
			fmt.Println(value)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Done go routine")
		done <- true
	}()
	fmt.Println("Series 2")
	<-done
}

func main() {
	name := "Florencia"
	fmt.Println("Synchronous")
	sync(name)
	fmt.Println("Concurrent Go routine")
	conc(name)
	fmt.Println("Concurrent Go routine and go channel")
	channel(name)
}
