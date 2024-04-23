package go_reloaded

import(
	"strings"
	"strconv"
)

func ApplyTransformations(wordList []string) []string {
	for i, word := range wordList {
		if word == "(up)" {
			wordList[i-1] = strings.ToUpper(wordList[i-1])
			wordList = append(wordList[:i], wordList[i+1:]...)
		} else if word == "(low)" {
			wordList[i-1] = strings.ToLower(wordList[i-1])
			wordList = append(wordList[:i], wordList[i+1:]...)
		} else if word == "(cap)" {
			wordList[i-1] = strings.Title(wordList[i-1])
			wordList = append(wordList[:i], wordList[i+1:]...)
		}else if word == "(up," {
			b := strings.Trim(string(wordList[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				wordList[i-j] = strings.ToUpper(wordList[i-j])
			}
			wordList = append(wordList[:i], wordList[i+2:]...)
			// Lower with number
		} else if word == "(low," {
			b := strings.Trim(string(wordList[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				wordList[i-j] = strings.ToLower(wordList[i-j])
			}
			wordList = append(wordList[:i], wordList[i+2:]...)
			// Capitalize with number
		} else if word == "(cap," {
			b := strings.Trim(string(wordList[i+1]), ")")
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				wordList[i-j] = strings.Title(wordList[i-j])
			}
			wordList = append(wordList[:i], wordList[i+2:]...)
		}
	}
	return wordList
}