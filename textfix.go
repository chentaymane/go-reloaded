package main

import (
	//"fmt"
	"strings"
)

func CheckPunc(s string) bool {
	countchek := 0
	if strings.HasPrefix(s, ".") || strings.HasPrefix(s, ",") || strings.HasPrefix(s, "!") || strings.HasPrefix(s, "?") || strings.HasPrefix(s, ":") || strings.HasPrefix(s, ";") {
		for _, k := range s {
			if k == '.' || k == ',' || k == '!' || k == '?' || k == ':' || k == ';' {
				countchek++
			}
		}
	}
	if len(s) == countchek {
		return true
	} else {
		return false
	}
}

func FixQuotes(lines []string) [][]string {
	words := []string{}
	temp := []string{}
	result := [][]string{}
	first := true
	str := ""
	for _, r := range lines {
		str = ""
		for _, k := range r {
			if k == '.' || k == ',' || k == '!' || k == '?' || k == ':' || k == ';' {
				str += " " + string(k) + " "
			} else {
				str += string(k)
			}
		}

		words = strings.Split(str, " ")

		words = Clean(words)
		words =AutoCorrect(words)
		for _, k := range words {

			if strings.Trim(k, "'") == "" {
				temp = append(temp, k)
				continue
			}
			if strings.HasSuffix(k, "'") && strings.HasPrefix(k, "'") && len(k) > 2 {
				temp = append(temp, k)
				continue
			}

			startQuotes := 0
			for strings.HasPrefix(k, "'") {
				startQuotes++
				k = k[1:]
			}

			if startQuotes > 0 {
				temp = append(temp, strings.Repeat("'", startQuotes))
			}

			finalQuotes := 0
			for strings.HasSuffix(k, "'") {
				finalQuotes++
				k = k[:len(k)-1]
			}

			if k != "" {
				temp = append(temp, k)
			}

			if finalQuotes > 0 {
				temp = append(temp, strings.Repeat("'", finalQuotes))
			}
		}

		temp = Clean(temp)
		temp = AtoAn(temp)
		// fmt.Println(temp)
		temp = AutoCorrect(temp)

		for i := 0; i < len(temp); i++ {

			r := temp[i]
			if i > 0 && (r == "." || r == "," || r == "!" || r == "?" || r == ":" || r == ";") {
				if temp[i-1] != "'" && !strings.Contains(temp[i-1], "'") {
					temp[i-1] = temp[i-1] + r
					temp[i] = ""
					temp = Clean(temp)
					i = 0
				}
			}
		}

		for i := 0; i < len(temp); i++ {

			r := temp[i]
			// Check if r contains only quotes
			if strings.Trim(r, "'") == "" && first && i+1 < len(temp) {
				temp[i+1] = r + temp[i+1]
				temp[i] = ""
				first = false

			} else if strings.Trim(r, "'") == "" && !first {
				temp[i-1] = temp[i-1] + r
				temp[i] = ""
				first = true
			}
		}
		first = true
		temp = Clean(temp)
		temp = AutoCorrect(temp)
		// fmt.Println(temp)

		for i := 0; i < len(temp); i++ {

			r := temp[i]
			if (i > 0 && (r == "." || r == "," || r == "!" || r == "?" || r == ":" || r == ";")) || i > 0 && CheckPunc(r) {
				temp[i-1] = temp[i-1] + r
				temp[i] = ""
				temp = Clean(temp)
				i = 0

			}
			temp = AutoCorrect(temp)
		}
		words = temp
		words = Clean(words)

		temp = []string{}
		words = AtoAn(words)
		result = append(result, words)

	}

	return result
}
