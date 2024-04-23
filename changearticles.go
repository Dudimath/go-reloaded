package go_reloaded

func ChangeArticles(words []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h"}

	for i, word := range words {
		for _, vowel := range vowels {
			if word == "a" && string(words[i+1][0]) == vowel {
				words[i] = "an"
			} else if word == "A" && string(words[i+1][0]) == vowel {
				words[i] = "An"
			}
		}
	}
	return words
}