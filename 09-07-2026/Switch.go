package main

import "time"

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Thursday:
		print("its weekend")
	default:
		print("its workday")
	}
}
