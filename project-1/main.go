package main

import (
	"fmt"
	"sync"
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
	var mtx sync.Mutex
	for i := 1; i < 5; i++ {
		mtx.Lock()
		wg.Add(1)
		go goroutine1(i, &wg)
		mtx.Unlock()
		mtx.Lock()
		wg.Add(1)
		go goroutine2(i, &wg)
		mtx.Unlock()
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

// func goroutine3(index int, wg *sync.WaitGroup) {

// 	fmt.Println([3]string{"coba1, coba2", "coba3"}, index)
// 	wg.Done()
// }

// func goroutine4(index int, wg *sync.WaitGroup) {
// 	fmt.Println([3]string{"bisa1, bisa2", "bisa3"}, index)
// 	wg.Done()
// }
