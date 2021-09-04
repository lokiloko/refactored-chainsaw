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

	// find the opening bracket
	indexFirstBracketFound := strings.Index(str, openBracketChar)

	// if no open bracket found means there will be no string inside bracket, thus return empty
	if indexFirstBracketFound >= 0 {
		runes := []rune(str)

		// collect all the words after first open bracket
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound + 1:len(str)])

		// find the closing bracket
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, closeBracketChar)

		// if no close bracket found after the opening bracket means there will be no string inside bracket
		// thus return empty
		if indexClosingBracketFound >= 0 {
			runes = []rune(wordsAfterFirstBracket)
			return string(runes[0 : indexClosingBracketFound])
		}
	}

	return ""
}
