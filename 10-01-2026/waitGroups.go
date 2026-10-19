package main

import (
	"fmt"
	"sync"
)

var sum1 int = 1

func task1(id int, wait_Group *sync.WaitGroup) {
	sum1 += 1
	defer wait_Group.Done()
	fmt.Println("doing task1 ", id)

}
func main() {
	// for i := 1; i <= 10; i++ {
	// 	task1(i)
	// }

	//here nouthing will print because main is also under thread
	//and it will end itself and end
	//so we will have to make sleep main function
	// for i := 1; i <= 10; i++ {
	// 	go task1(i)
	// }

	var wait_Group sync.WaitGroup
	//all task1 will run concurrently and very fast
	for i := 1; i <= 1000; i++ {
		wait_Group.Add(1)
		go task1(i, &wait_Group)
	}
	wait_Group.Wait()
	fmt.Println(sum1)

}
