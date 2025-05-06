package main

import (
	"sort"
)

// get the frequency map
func GetFrequencyMap(arr []int) map[int]int {
	frequency := make(map[int]int)
	for _, v := range arr {
		frequency[v]++
	}
	return frequency
}

// sort map keys depending on val ues
func sortMapByValues(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}

// take top k elements
func TopKFrequent(nums []int, k int) []int {
	m := GetFrequencyMap(nums)
	keys := sortMapByValues(m)
	ans := make([]int, 0)
	for i, key := range keys {
		if i+1 <= k {
			ans = append(ans, key)
		}
	}
	return ans
}
