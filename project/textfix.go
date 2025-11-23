package main

import (
	"strings"
)

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
				str += " "+string(k) + " "
			} else {
				str += string(k)
			}
		}

		words = strings.Split(str, " ")
		words = Clean(words)

		for _, k := range words {
			if strings.Trim(k, "'") == "" {
				temp = append(temp, k)
				continue
			}

			leadingQuotes := 0
			for strings.HasPrefix(k, "'") {
				leadingQuotes++
				k = k[1:]
			}
			if leadingQuotes > 0 {
				temp = append(temp, strings.Repeat("'", leadingQuotes))
			}

			trailingQuotes := 0
			for strings.HasSuffix(k, "'") {
				trailingQuotes++
				k = k[:len(k)-1]
			}

			if k != "" {
				temp = append(temp, k)
			}

			if trailingQuotes > 0 {
				temp = append(temp, strings.Repeat("'", trailingQuotes))
			}
		}

		temp = Clean(temp)
		temp = AtoAn(temp)

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

		temp = AutoCorrect(temp)

		for i := 0; i < len(temp); i++ {
			r := temp[i]
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

		for i := 0; i < len(temp); i++ {
			r := temp[i]
			if i > 0 && (r == "." || r == "," || r == "!" || r == "?" || r == ":" || r == ";") {
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
