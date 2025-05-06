package main

import (
	"fmt"
	"slices"
	"strconv"
)

// not accepted
func threeSumBruteForce(nums []int) [][]int {
	n := len(nums)
	m := make(map[string][]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					store := []int{nums[i], nums[j], nums[k]}
					slices.Sort(store)
					str := ""
					for _, v := range store {
						str += strconv.Itoa(v)
					}
					m[str] = store
				}
			}
		}
	}
	ans := [][]int{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	m := make(map[string][]int)
	slices.Sort(nums)
	for i := 0; i < n; i++ {
		a := nums[i]
		ptr1, ptr2 := 0, n-1
		for ptr1 < ptr2 {
			if i != ptr1 && i != ptr2 {
				sum := nums[ptr1] + a + nums[ptr2]
				if sum < 0 {
					ptr1++
				} else if sum > 0 {
					ptr2--
				} else {
					store := []int{nums[ptr1], nums[ptr2], a}
					slices.Sort(store)
					str := ""
					for _, v := range store {
						str += strconv.Itoa(v)
					}
					m[str] = store
				}
			}
		}
	}
	ans := [][]int{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
