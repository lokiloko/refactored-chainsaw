package main

import (
	"fmt"
	"strings"
)

func main() {
	// Driver Code
	testCase1 := "(abcdefghijklmn)"
	testCase2 := "Saya suka coding (IYKWIM)"
	testCase3 := "Kalimat dengan kurung buka ( tanpa kurung tutup"
	testCase4 := "Statement dengan 2 tanda (ini satu) dan (ini yg kedua)"

	fmt.Println(findFirstStringInBracket(testCase1))
	fmt.Println(findFirstStringInBracket(testCase2))
	fmt.Println(findFirstStringInBracket(testCase3))
	fmt.Println(findFirstStringInBracket(testCase4))
}

func findFirstStringInBracket(str string) string {
	openBracketChar := "("
	closeBracketChar := ")"

	indexFirstBracketFound := strings.Index(str, openBracketChar)
	if indexFirstBracketFound >= 0 {
		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, closeBracketChar)
		if indexClosingBracketFound >= 0 {
			runes := []rune(wordsAfterFirstBracket)
			return string(runes[1 : indexClosingBracketFound-1])
		}
	}

	// if no open bracket found means there will be no string inside bracket, thus return empty
	return ""
}
