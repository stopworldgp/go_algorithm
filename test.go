package main

import "fmt"

func t(nums []int) {
	fmt.Println(nums[1])
}

func main() {

	nums := []int{1, 2}
	fmt.Println(nums[1:1])
}
