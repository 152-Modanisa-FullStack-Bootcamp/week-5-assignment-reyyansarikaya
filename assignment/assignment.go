package assignment

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := uint64(x) + uint64(y)
	overflow := uint32(sum>>32) == 1
	return x + y, overflow
}

func CeilNumber(f float64) float64 {
	number, frac := math.Modf(f)
	if 0 == frac {
		return f
	}
	if 0 < frac && frac <= 0.25 {
		return number + 0.25
	}

	if 0.25 < frac && frac <= 0.50 {
		return number + 0.50
	}

	if 0.50 < frac && frac <= 0.75 {
		return number + 0.75
	}
	if 0.75 < frac && float32(frac) <= 0.99 {
		return math.Ceil(f)
	}
	return f
}

func AlphabetSoup(s string) string {
	stringArray := strings.Split(s, "")
	sort.Strings(stringArray)
	justString := strings.Join(stringArray, "")
	return justString
}

func StringMask(s string, n uint) string {
	stringArray := strings.Split(s, "")
	if len(stringArray) <= 1 {
		return "*"
	}
	if int(n) >= len(stringArray) {
		n = 0
	}
	for j := int(n); j < len(stringArray); j++ {
		stringArray[j] = "*"
	}
	justString := strings.Join(stringArray, "")
	return justString
}

// For Benchmark Test
func StringMask_SadumanSolve(s string, n uint) string {
	if len(s) == 0 {
		return "*"
	} else if len(s) <= int(n) {
		return strings.Repeat("*", len(s))
	} else {
		return s[:n] + strings.Repeat("*", len(s)-int(n))
	}
}

func WordSplit(arr [2]string) string {
	index := 0
	word := ""
	matchedWord := ""
	wordsArray := strings.Split(arr[1], ",")
	wordArray := strings.Split(arr[0], "")
out:
	for i := 0; i < len(wordArray); i++ {
		for j := 0; j < len(wordsArray); j++ {
			if index < len([]rune(arr[0])) {
				word += wordArray[index]
				index++
			}
			matched, _ := regexp.MatchString(wordsArray[j], word)
			if matched && word == wordsArray[j] {
				matchedWord += word
				word = ""
				if len([]rune(arr[0]))+1 != len([]rune(matchedWord)) {
					matchedWord += ","
				} else {
					break out
				}
			}
		}
	}
	if len([]rune(arr[0]))+1 == len([]rune(matchedWord)) {
		return matchedWord
	}
	return "not possible"
}

func VariadicSet(i ...interface{}) []interface{} {
	// 🎶 💃🕺
	checkMap := make(map[interface{}]bool)
	returnList := []interface{}{}
	for _, interfaceValue := range i {
		_, isThere := checkMap[interfaceValue]
		if !isThere {
			checkMap[interfaceValue] = true
			returnList = append(returnList, interfaceValue)
		}
	}
	return returnList
}
