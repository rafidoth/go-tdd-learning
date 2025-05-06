package main

import "fmt"

func SuffixPrefixWay(nums []int) []int {
	n := len(nums)
	suffix, prefix := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		prefix[i] = 1
		suffix[i] = 1
	}
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i]
	}
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		prod := 1
		if i-1 >= 0 {
			prod *= prefix[i-1]
		}
		if i+1 < n {
			prod *= suffix[i+1]
		}
		nums[i] = prod
	}
	return nums
}

func SuffixPrefixWaySpaceOptimized(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix * nums[i]
		prefix = answer[i]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		suffix *= nums[i]
		answer[i] = suffix
		if i-1 >= 0 {
			nums[i] = answer[i-1]
		} else {
			nums[i] = 1
		}
	}

	for i := 0; i < n-1; i++ {
		nums[i] *= answer[i+1]
	}
	return nums
}
func productExceptSelf(nums []int) []int {
	// return SuffixPrefixWay(nums)
	return SuffixPrefixWaySpaceOptimized(nums)
}

func main() {
	nums := []int{2, 1, 4, 3}
	fmt.Println(productExceptSelf(nums))
}
