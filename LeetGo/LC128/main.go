package main

func countConsecutive(m map[int]bool, start int) int {
	count := 0
	for {
		if m[start] {
			count++
			start++
		} else {
			break
		}
	}
	return count
}

func longestConsecutive(nums []int) int {
	isP := make(map[int]bool)

	for _, val := range nums {
		isP[val] = true
	}

	maxCount := 0
	for key := range isP {
		if !isP[key-1] {
			c := countConsecutive(isP, key)
			if c > maxCount {
				maxCount = c
			}
		}
	}
	return maxCount
}

func main() {

}
