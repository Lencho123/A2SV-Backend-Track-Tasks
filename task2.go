package main

import (
	"fmt"
	"strings"
	"unicode"
)

func WordFrequency(input string) map[string]int {

	// Convert the input string to lowercase
	input = strings.ToLower(input)

	// Remove punctuation from the input string

	clean_input := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, input)

	// Split the cleaned input into words
	words := strings.Fields(clean_input)

	// Count the frequency of each word
	word_freq := make(map[string]int)
	for _, word := range words {
		word_freq[word]++
	}

	return word_freq
}

func IsPalindrome(input string) bool {
	// Convert input to list of lowercase letters and digits
	var chars []rune
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			chars = append(chars, unicode.ToLower(char))
		}
	}

	// Check if the cleaned characters form a palindrome
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}

func main() {

	text := "Hello, world! Hello, universe. Hello world."

	frequency := WordFrequency(text)

	fmt.Println(frequency)

	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("Hello, world!"))
}
