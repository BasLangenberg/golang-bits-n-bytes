package main

import (
	"fmt"
	"sync"
	"time"
)

func method(wg *sync.WaitGroup, input string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v round %d\n", input, i)
		time.Sleep(time.Millisecond)
	}
}

func receiver(wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()

	i := 0
	for s := range channel {
		fmt.Println(s)

		i++
		if i >= 20 {
			break
		}
	}
}

func sender(wg *sync.WaitGroup, input string, channel chan string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%v round %d", input, i)
		channel <- s
		time.Sleep(time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go method(&wg, "goroutine 1")
	go method(&wg, "goroutine 2")

	printChannel := make(chan string)

	go receiver(&wg, printChannel)

	go sender(&wg, "goroutine 3", printChannel)
	go sender(&wg, "goroutine 4", printChannel)

	wg.Wait()
}
