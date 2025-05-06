package group_anagrams

import (
	"maps"
	"slices"
	"strconv"
)

func getFrequencyString(str string) string {
	frequency := make(map[rune]int)
	for _, ch := range str {
		frequency[ch]++
	}
	// Sort the frequency map by key
	keys := slices.Sorted(maps.Keys(frequency))

	freqString := ""
	for _, key := range keys {
		freqString += string(key)
		freqString += strconv.Itoa(frequency[key])
	}
	return freqString
}

func groupAnagrams1(strs []string) [][]string {
	dictionary := make(map[string][]string)

	for _, str := range strs {
		fstr := getFrequencyString(str)
		dictionary[fstr] = append(dictionary[fstr], str)
	}

	ans := [][]string{}
	for _, record := range dictionary {
		ans = append(ans, record)
	}
	return ans
}
