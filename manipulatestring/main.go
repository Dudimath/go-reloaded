package main

import (
	"os"
	"strings"
	"go_reloaded"
)

func main() {
	args := os.Args[1:]

	// Reading  file

	content, _ := os.ReadFile(args[0])

	// Array to hold the words

	wordList := strings.Split(string(content), " ")
	wordList = go_reloaded.ApplyTransformations(wordList)
	for i, word := range wordList {
		if word == "(bin)" {
			wordList[i-1] = go_reloaded.BintoInt(wordList[i-1])
			wordList = append(wordList[:i],wordList[i+1:]... )
		} else if word == "(hex)" {
			wordList[i-1] = go_reloaded.HextoInt(wordList[i-1])
			wordList = append(wordList[:i],wordList[i+1:]... )
		}
	}
	go_reloaded.ChangeArticles(wordList)

	// Joining the slice
	result := strings.Join(go_reloaded.Punctuations(wordList), " ") + "\n"

	// Write the manipulated file
	err := os.WriteFile(args[1], []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}




