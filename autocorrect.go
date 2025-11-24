package main

import (
	"strconv"
	"strings"
)

func AutoCorrect(words []string) []string {
	for i := 0; i < len(words); i++ {
		r := words[i]

		if r == "(hex)" {
			if i > 0 {
				decimal := Hex(words[i-1])
				if decimal != -1 {
					words[i-1] = strconv.Itoa(decimal)
				}
			}

			words[i] = ""
			words = Clean(words)

		}

		if r == "(bin)" {
			if i > 0 {
				decimal := Bin(words[i-1])
				if decimal != -1 {
					words[i-1] = strconv.Itoa(decimal)
				}
			}
			words[i] = ""
			words = Clean(words)

		}

		if r == "(cap)" {
			if i > 0 {
				words[i-1] = Cap(words[i-1])
			}

			words[i] = ""
			words = Clean(words)

		}

		if r == "(cap" && words[i+1] == "," && i+2 < len(words) {
			valueStr := words[i+2]
			if strings.HasSuffix(valueStr, ")") {
				valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
				if i >= valueInt {
					for j := i - 1; j >= i-valueInt; j-- {
						words[j] = Cap(words[j])
					}
				}
				words[i] = ""
				words[i+1] = ""
				words[i+2] = ""
			}

			words = Clean(words)
		}

		if r == "(low)" {
			if i > 0 {
				words[i-1] = Low(words[i-1])
			}

			words[i] = ""
			words = Clean(words)

		}

		if r == "(low" && words[i+1] == "," && i+2 < len(words) {
			valueStr := words[i+2]
			if strings.HasSuffix(valueStr, ")") {
				valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
				if i >= valueInt {
					for j := i - 1; j >= i-valueInt; j-- {
						words[j] = Low(words[j])
					}
				}
				words[i] = ""
				words[i+1] = ""
				words[i+2] = ""

			}
			words = Clean(words)
		}

		if r == "(up)" {
			if i-1 >= 0 {
				words[i-1] = Up(words[i-1])
			}

			words[i] = ""
			words = Clean(words)

		}

		if r == "(up" && words[i+1] == "," && i+2 < len(words) {
			valueStr := words[i+2]
			if strings.HasSuffix(valueStr, ")") {

				valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
				if i >= valueInt {
					for j := i - 1; j >= i-valueInt; j-- {
						words[j] = Up(words[j])
					}
				}
				words[i] = ""
				words[i+1] = ""
				words[i+2] = ""

			}

			words = Clean(words)
		}
	}

	return words
}
