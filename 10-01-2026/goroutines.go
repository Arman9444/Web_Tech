package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task ", id)
}
func main() {
	// for i := 1; i <= 10; i++ {
	// 	task(i)
	// }

	//here nouthing will print because main is also under thread
	//and it will end itself and end
	//so we will have to make sleep main function
	// for i := 1; i <= 10; i++ {
	// 	go task(i)
	// }

	//all task will run concurrently and very fast
	for i := 1; i <= 10; i++ {
		go task(i)
	}
	time.Sleep(time.Millisecond)

}
