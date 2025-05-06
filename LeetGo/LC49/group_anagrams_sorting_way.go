package group_anagrams

import (
	"sort"
)

// sorting takes around O(nlogn) time
// using the most efficient algorithm
func SortStr(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	dictionary := make(map[string][]string)

	for _, str := range strs {
		sortedStr := SortStr(str)
		dictionary[sortedStr] = append(dictionary[sortedStr], str)
	}
	ans := [][]string{}
	for _, record := range dictionary {
		ans = append(ans, record)
	}
	return ans
}
