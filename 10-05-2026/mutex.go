package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}
func main() {
	var waitGroup sync.WaitGroup
	myPost := post{
		views: 0,
	}
	for i := 1; i <= 1000; i++ {
		waitGroup.Add(1)
		go myPost.increment(&waitGroup)
	}
	// myPost.increment()
	// myPost.increment()
	waitGroup.Wait()
	fmt.Println(myPost.views)
}
