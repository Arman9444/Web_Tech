// constant
package main

import "fmt"

const age = 30

var p = 29

func main() {
	const name = "arman"

	//name="dep"//error-> can't reassign
	fmt.Println(age)
	const (
		port = 5000
		host = "localHost"
	)
	fmt.Println(port, host)
	println(age)

}
