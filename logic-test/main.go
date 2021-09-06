package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Driver Code
	testCase1 := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	testCase2 := []string{"bakar", "kabar", "sabar", "baka", "kaba", "kbaa", "a"}
	testCase3 := []string{"abcdefg", "bcdefga", "cdefgab", "defgabc", "efgabcd", "fgabcde", "gabcdef"}

	fmt.Println(solveAnagram(testCase1))
	fmt.Println(solveAnagram(testCase2))
	fmt.Println(solveAnagram(testCase3))
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