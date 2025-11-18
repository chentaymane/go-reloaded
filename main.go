package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Hex(arg string) int {
	value := arg
	result, _ := strconv.ParseInt(value, 16, 64)
	return int(result)
}

func Bin(arg string) int {
	value := arg
	result, _ := strconv.ParseInt(value, 2, 64)
	return int(result)
}

func Cap(s string) string {
	runes := []rune(s)
	isNewWord := true

	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9') {
			if isNewWord {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] = runes[i] - 'a' + 'A'
				}
				isNewWord = false
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] - 'A' + 'a'
				}
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

func main() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

words := []string{}

for _, line := range lines {
    current := ""
    for _, k := range line {
        if k != ' ' && k != '\t' {
            current += string(k)
        } else {
            if current != "" {
                words = append(words, current)
                current = ""
            }
        }
		
    }
	
  if current != "" {
        words = append(words, current)
    }
	if len(words)!= 0 {
			words = append(words, "\n")
		}
}

				fmt.Println(words)
	for i, r := range words {
		if r == "(hex)" {
			word := Hex(string(words[i-1]))
			newWord := strconv.Itoa(word)
			words[i-1] = newWord
			words[i] = ""
			content = []byte(strings.Join(words, " "))
		}
		if r == "(bin)" {
			word := Bin(string(words[i-1]))
			newWord := strconv.Itoa(word)
			words[i-1] = newWord
			words[i] = ""
			content = []byte(strings.Join(words, " "))
		}
		if r == "(cap)" {
			word := Cap(words[i-1])
			words[i-1] = word
			words[i] = ""
			content = []byte(strings.Join(words, " "))
		}
		if r == "(cap," {
			valueStr := words[i+1]                                  
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1]) 

			for j := i - 1; j >= i-valueInt; j-- {
				words[j] = Cap(words[j])
			}

			words[i] = ""   
			words[i+1] = "" 

			content = []byte(strings.Join(words, " "))
		}

		if r == "(low)" {
			word := Low(words[i-1])
			words[i-1] = word
			words[i] = ""
			content = []byte(strings.Join(words, " "))
		}
		if r == "(low," {
			valueStr := words[i+1]                                  
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1]) 

			for j := i - 1; j >= i-valueInt; j-- {
				words[j] = Low(words[j])
			}

			words[i] = ""   
			words[i+1] = "" 

			content = []byte(strings.Join(words, " "))
		}
		if r == "(up)" {
			word := Up(words[i-1])
			words[i-1] = word
			words[i] = ""
			content = []byte(strings.Join(words, " "))
		}
		if r == "(up," {
			valueStr := words[i+1]                                  
			valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1]) 

			for j := i - 1; j >= i-valueInt; j-- {
				words[j] = Up(words[j])
			}

			words[i] = ""   
			words[i+1] = "" 

			content = []byte(strings.Join(words, " "))
		}

	}

	fmt.Println(string(content))
}
