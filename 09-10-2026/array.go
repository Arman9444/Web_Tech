package main

import "fmt"

func main() {
	var x = 9
	var nums [6]int
	nums[0] = 1
	//array length
	fmt.Println(len(nums))
	//print array
	fmt.Println(nums)
	println(nums[1], x)
	var vals [5]bool
	vals[3] = true
	fmt.Println(vals)
	var str [4]string
	str[1] = "golang"
	fmt.Println(str)

	nums1 := [4]int{1, 4, 5, 0}
	fmt.Println(nums1)

	nums2 := [4]int{}
	fmt.Println(nums2)
	//2d array
	var nums3 [2][2]int
	fmt.Println(nums3)

	nums4 := [2][3]int{{5, 8, 9}, {8, 9, 2}}
	fmt.Println(nums4)

}
