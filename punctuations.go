package go_reloaded


func Punctuations(words []string) []string {
	punctuations := []string{",", ".", "!", "?", ":", ";"}
	// Punctuation in the middle of a string connecting to the word after
	for i, word := range words {
		for _, punctuation := range punctuations {
			if string(word[0]) == punctuation && string(word[len(word)-1]) != punctuation {
				words[i-1] += punctuation
				words[i] = word[1:]
			}
		}
	}

	// Punctuation at the end of the string
	for i, word := range words {
		for _, punctuation := range punctuations {
			if string(word[0]) == punctuation && words[len(words)-1] == words[i] {
				words[i-1] += word
				words = words[:len(words)-1]
			}
		}
	}

	// Punctuation in the middle of the string
	for i, word := range words {
		for _, punctuation := range punctuations {
			if string(word[0]) == punctuation && string(word[len(word)-1]) == punctuation && words[i] != words[len(words)-1] {
				words[i-1] += word
				words = append(words[:i], words[i+1:]...)
			}
		}
	}

	// For apostrophe
	count := 0
	for i, word := range words {
		if word == "'" && count == 0 {
			count++
			words[i+1] = word + words[i+1]
			words = append(words[:i], words[i+1:]...)
		}
	}
	// For the second apostrophe
	for i, word := range words {
		if word == "'" {
			words[i-1] += word
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
