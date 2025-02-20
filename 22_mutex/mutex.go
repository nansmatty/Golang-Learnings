package main

import (
	"fmt"
	"sync"
)

// Race Condition :== When multiple goroutines single resource or access/modify a varaible without synchronization. This leads to unpredictable behavior, such as data corruption or inconsistent results.

// Good Practice is use the mutex Lock() only the place where the modification is happening otherwise this will lock the entire func wil be lock this will create a small bottleneck....

type post struct {
	views int
	mutex sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mutex.Unlock()
		wg.Done()
	}()

	p.mutex.Lock()
	p.views += 1
}

func main() {
	fmt.Println("Mutex")

	var wg sync.WaitGroup

	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
