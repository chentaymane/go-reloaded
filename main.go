package main

import (
	"fmt"
	"os"
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
		result = result + string(r)
	}
	return result
}

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

		}

		if r == "(bin)" {
			if i > 0 {
				decimal := Bin(words[i-1])
				if decimal != -1 {
					words[i-1] = strconv.Itoa(decimal)
				}
			}
			words[i] = ""

		}

		if r == "(cap)" {
			if i > 0 {
				words[i-1] = Cap(words[i-1])
			}

			words[i] = ""

		}

		if r == "(cap," {

			valueStr := words[i+1]
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
			if i >= valueInt {
				for j := i - 1; j >= i-valueInt; j-- {
					words[j] = Cap(words[j])
				}
			}

			words[i] = ""
			words[i+1] = ""

		}

		if r == "(low)" {
			if i > 0 {
				words[i-1] = Low(words[i-1])
			}

			words[i] = ""

		}

		if r == "(low," {
			valueStr := words[i+1]
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
			if i >= valueInt {
				for j := i - 1; j >= i-valueInt; j-- {
					words[j] = Low(words[j])
				}
			}
			words[i] = ""
			words[i+1] = ""

		}

		if r == "(up)" {
			if i-1 > 0 {
				words[i-1] = Up(words[i-1])
			}
			words[i] = ""

		}

		if r == "(up," {
			valueStr := words[i+1]
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
			if i >= valueInt {
				for j := i - 1; j >= i-valueInt; j-- {
					words[j] = Up(words[j])
				}
			}
			words[i] = ""
			words[i+1] = ""

		}
	}

	return words
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
			if (r == "a" || r == "A") && (strings.HasPrefix(Low(words[i+1]), "a") || strings.HasPrefix(Low(words[i+1]), "e") || strings.HasPrefix(Low(words[i+1]), "i") || strings.HasPrefix(Low(words[i+1]), "o") || strings.HasPrefix(Low(words[i+1]), "u") || strings.HasPrefix(Low(words[i+1]), "h")) {
				words[i] = words[i] + "n"
			}
		}
	}
	return words
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
				str += string(k) + " "
			} else {
				str += string(k)
			}

		}

		words = strings.Split(str, " ")
		words = Clean(words)
		for i, k := range words {

			if strings.HasPrefix(k, "'") && !strings.HasSuffix(k, "'") {
				temp = append(temp, "'")
				temp = append(temp, words[i][1:])
			} else if strings.HasSuffix(k, "'") && !strings.HasPrefix(k, "'") {
				temp = append(temp, words[i][:len(k)-1])
				temp = append(temp, "'")

			} else {
				temp = append(temp, k)

			}
		}
		for i := 0; i < len(temp); i++ {
			r := temp[i]
			if r == "'" && first && i+1 < len(temp) {
				temp[i+1] = "'" + temp[i+1]
				temp[i] = ""
				first = false
				continue
			} else if r == "'" {
				temp[i-1] = temp[i-1] + "'"
				temp[i] = ""
				first = true
			}

		}
		first = true

		words = temp
		words = Clean(words)
		for i := 0; i < len(words); i++ {

			r := words[i]
			if r == "." || r == "," || r == "!" || r == "?" || r == ":" || r == ";" {
				words[i-1] = words[i-1] + r
				words[i] = ""
				words = Clean(words)
				i = 0

			}
			words = AutoCorrect(words)
		}
		temp = []string{}
		words = AtoAn(words)
		result = append(result, words)

	}

	return result

}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Check file extensions
	if !strings.HasSuffix(strings.ToLower(inputFile), ".txt") {
		fmt.Println("Error: input file must have .txt extension")
		return
	}

	if !strings.HasSuffix(strings.ToLower(outputFile), ".txt") {
		fmt.Println("Error: output file must have .txt extension")
		return
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	result := FixQuotes(lines)
	resultF := ""

	for i, r := range result {
		for _, k := range Clean(r) {
			resultF += k + " "
		}
		if i < len(result)-1 {
			resultF += "\n"
		}
	}

	err = os.WriteFile(outputFile, []byte(resultF), 0644)
	if err != nil {
		fmt.Println("Error writing output:", err)
		return
	}
}
