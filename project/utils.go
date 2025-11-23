package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Hex(arg string) int {
	value := arg
	result, err := strconv.ParseInt(value, 16, 64)
	if err != nil {
		return -1
	}
	return int(result)
}

func Bin(arg string) int {
	value := arg
	result, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		return -1
	}
	return int(result)
}

func Cap(s string) string {
	runes := []rune(s)
	isNewWord := true

	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]) {
			if isNewWord {
				runes[i] = unicode.ToUpper(runes[i])
				isNewWord = false
			} else {
				runes[i] = unicode.ToLower(runes[i])
			}
		} else {
			isNewWord = true
		}
	}

	return string(runes)
}

func Low(s string) string {
	var result string
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		result += string(r)
	}
	return result
}

func Up(s string) string {
	var result string
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		result += string(r)
	}
	return result
}

func Clean(words []string) []string {
	wordsClean := []string{}
	for _, r := range words {
		if r != "" {
			wordsClean = append(wordsClean, r)
		}
	}
	return wordsClean
}

func AtoAn(words []string) []string {
	for i, r := range words {
		if i+1 < len(words) {
			if (r == "a" || r == "A") && (strings.HasPrefix(Low(words[i+1]), "a") ||
				strings.HasPrefix(Low(words[i+1]), "e") ||
				strings.HasPrefix(Low(words[i+1]), "i") ||
				strings.HasPrefix(Low(words[i+1]), "o") ||
				strings.HasPrefix(Low(words[i+1]), "u") ||
				strings.HasPrefix(Low(words[i+1]), "h")) {
				words[i] = words[i] + "n"
			}
		}
	}
	return words
}
