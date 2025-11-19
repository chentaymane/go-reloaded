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
	str := ""

	for i, r := range lines {
		for _, k := range r {
			if k == '.' || k == ',' || k == '!' || k == '?' || k == ':' || k == ';' {
				str += string(k) + " "
			} else {
				str += string(k)
			}
		}
		if r != "" && i < len(lines)-1 {
			str += " \n"
		}
	}

	tab := strings.Split(str, " ")
	words := []string{}

	for _, r := range tab {
		if r != "" {
			words = append(words, r)
		}
	}

	for i, r := range words {
		if i+1 < len(words) {
			if (r == "a" || r == "A") && (strings.HasPrefix(Low(words[i+1]), "a") || strings.HasPrefix(Low(words[i+1]), "e") || strings.HasPrefix(Low(words[i+1]), "i") || strings.HasPrefix(Low(words[i+1]), "o") || strings.HasPrefix(Low(words[i+1]), "u") || strings.HasPrefix(Low(words[i+1]), "h")) {
				words[i] = words[i] + "n"
			}
		}
	}

	fmt.Println(words)
	fmt.Println()
	for i := 0; i < len(words); i++ {
    r := words[i]

    // ---------------- HEX ------------------
    if r == "(hex)" && i > 0 {
        decimal := Hex(words[i-1])
        words[i-1] = strconv.Itoa(decimal)
        words[i] = ""
    }

    if strings.HasSuffix(r, "(hex)") {
        base := strings.TrimSuffix(r, "(hex)")
        decimal := Hex(base)
        words[i] = strconv.Itoa(decimal)
    }

    // ---------------- BIN ------------------
    if r == "(bin)" && i > 0 {
        decimal := Bin(words[i-1])
        words[i-1] = strconv.Itoa(decimal)
        words[i] = ""
    }

    if strings.HasSuffix(r, "(bin)") {
        base := strings.TrimSuffix(r, "(bin)")
        decimal := Bin(base)
        words[i] = strconv.Itoa(decimal)
    }

    // ---------------- CAP ------------------
    if r == "(cap)" && i > 0 {
        words[i-1] = Cap(words[i-1])
        words[i] = ""
    }

    if strings.HasSuffix(r, "(cap)") {
        base := strings.TrimSuffix(r, "(cap)")
        words[i] = Cap(base)
    }

    // cap with number: (cap, X)
    if r == "(cap," {
        valueStr := words[i+1]
        valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
        for j := i - 1; j >= i-valueInt; j-- {
            words[j] = Cap(words[j])
        }
        words[i] = ""
        words[i+1] = ""
    }

    // ---------------- LOW ------------------
    if r == "(low)" && i > 0 {
        words[i-1] = Low(words[i-1])
        words[i] = ""
    }

    if strings.HasSuffix(r, "(low)") {
        base := strings.TrimSuffix(r, "(low)")
        words[i] = Low(base)
    }

    // low with number: (low, X)
    if r == "(low," {
        valueStr := words[i+1]
        valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
        for j := i - 1; j >= i-valueInt; j-- {
            words[j] = Low(words[j])
        }
        words[i] = ""
        words[i+1] = ""
    }

    // ---------------- UP ------------------
    if r == "(up)" && i > 0 {
        words[i-1] = Up(words[i-1])
        words[i] = ""
    }

    if strings.HasSuffix(r, "(up)") {
        base := strings.TrimSuffix(r, "(up)")
        words[i] = Up(base)
    }

    // up with number: (up, X)
    if r == "(up," {
        valueStr := words[i+1]
        valueInt, _ := strconv.Atoi(valueStr[:len(valueStr)-1])
        for j := i - 1; j >= i-valueInt; j-- {
            words[j] = Up(words[j])
        }
        words[i] = ""
        words[i+1] = ""
    }
}

	wordsClean := []string{}
	for _, r := range words {
		if r != "" {
			wordsClean = append(wordsClean, r)
		}
	}
	words = wordsClean

	for i := 0; i < len(words); i++ {
		r := words[i]
		if r == "." || r == "," || r == "!" || r == "?" || r == ":" || r == ";" {
			words[i-1] = words[i-1] + r
			words[i] = ""
			wordsF := []string{}
			for _, r := range words {
				if r != "" {
					wordsF = append(wordsF, r)
				}
			}
			words = wordsF
			i--
		}
	}

	for i := 0; i < len(words); i++ {
		r := words[i]
		if r == "'" {
			if i+1 < len(words) {
				words[i+1] = r + words[i+1]
				words[i] = ""
				wordsF := []string{}
				for _, w := range words {
					if w != "" {
						wordsF = append(wordsF, w)
					}
				}
				words = wordsF
				i--
			}
		} else if strings.HasSuffix(r, "'") && len(r) > 1 {
			continue
		} else if strings.HasPrefix(r, "'") {
			words[i-1] = words[i-1] + r
			words[i] = ""
			wordsF := []string{}
			for _, w := range words {
				if w != "" {
					wordsF = append(wordsF, w)
				}
			}
			words = wordsF
			i--
		}
	}

	content = []byte(strings.Join(words, " "))
	fmt.Println(string(content))
}
