package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("Starting execution in: F1")
	for i := 0; i < 5; i++ {
		fmt.Println("In F1, the value of X is:", i)
		time.Sleep(time.Second)

	}
	fmt.Println("Done executing: F1")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("Starting execution in: F2")
	for i := 0; i < 10; i++ {
		fmt.Println("In F2, the value of Y is:", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Done executing: F2")
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	fmt.Println("Main thread executing")
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Number of CPUS: ", runtime.NumCPU())
	fmt.Println("Arch ", runtime.GOARCH)

	
	var n = 0

	var m sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func() {
			m.Lock()
			n++
			m.Unlock()
			time.Sleep(time.Second / 10)
			defer wg.Done()

		}()
		go func() {
			m.Lock()
			n--
			m.Unlock()

			time.Sleep(time.Second / 10)
			defer wg.Done()

		}()
	}

	defer fmt.Println("N =", n)

	wg.Wait()

	fmt.Println("Go max processes: ", runtime.GOMAXPROCS(0))

}
