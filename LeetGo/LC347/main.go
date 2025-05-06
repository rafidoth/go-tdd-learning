package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	ans := TopKFrequent2(nums, 2)
	fmt.Println(ans)
}
