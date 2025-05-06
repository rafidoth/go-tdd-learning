package main

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

func TopKFrequent2(nums []int, k int) []int {
	m := GetFrequencyMap(nums)
	// this is a min-heap
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		aAsserted := a.([2]int)
		bAsserted := b.([2]int)
		return utils.IntComparator(aAsserted[1], bAsserted[1])
	})
	for n, freq := range m {
		pq.Enqueue([2]int{n, freq})
		if pq.Size() > k {
			pq.Dequeue()
		}
	}
	ans := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result, _ := pq.Dequeue()
		ans[i] = result.([2]int)[0]
	}
	return ans
}
