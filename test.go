package main

import (
	"fmt"
	"math/rand"
	"time"
)

func t(nums []int) {
	fmt.Println(nums[1])
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))
}
