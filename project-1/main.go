package main

import (
	"fmt"
	"sync"
)

var (
	mtx sync.Mutex
)

func main() {
	// GOROUTINE keduanya menampilkan secara acak
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(2)
		go goroutine1(i, &wg)
		go goroutine2(i, &wg)
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("===================")
	fmt.Println()

	// GOROUTINE keduanya menampilkan secara rapih
	// var mtx sync.Mutex
	for i := 1; i < 5; i++ {
		wg.Add(1)
		mtx.Lock()
		go goroutine1(i, &wg)
		mtx.Unlock()
		wg.Wait()
		wg.Add(1)

		mtx.Lock()
		go goroutine2(i, &wg)
		mtx.Unlock()
		wg.Wait()
	}

	wg.Wait()

}

func goroutine1(index int, wg *sync.WaitGroup) {
	fmt.Println([3]string{"coba1, coba2", "coba3"}, index)
	wg.Done()
}

func goroutine2(index int, wg *sync.WaitGroup) {
	fmt.Println([3]string{"bisa1, bisa2", "bisa3"}, index)
	wg.Done()
}
