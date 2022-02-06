package main

import (
	"fmt"
	"sync"
)

func SyncWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printString("Salam", &wg)
	go printString("Hello", &wg)

	wg.Wait()
}

func printString(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}
