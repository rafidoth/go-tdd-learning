package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(groups [][]string) [][]string {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		// Handle empty groups just in case
		if len(groups[i]) == 0 || len(groups[j]) == 0 {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
	return groups
}
func TestGroupAnagrams(t *testing.T) {
	input1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected_output1 := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}

	solutions := []func([]string) [][]string{groupAnagrams, groupAnagrams1}

	for _, solution := range solutions {
		res := solution(input1)
		if !reflect.DeepEqual(normalize(res), normalize(expected_output1)) {
			t.Error("Failed, expected:", expected_output1, "but got:", res)
		}
	}
}
