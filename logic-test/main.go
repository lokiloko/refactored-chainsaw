package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Driver Code
	testCase1 := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

	fmt.Println(solveAnagram(testCase1))
}

func solveAnagram(testCase []string) [][]string {
	result := [][]string{}
	anagrams := map[string][]string{}

	for _, str := range testCase {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	for _, anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}