package problems

func reverseWords(words []rune) {

	// reverse whole string
	// "hello world in go" -> "og ni dlrow olleh"
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}


	// find the end of each world and reverse it
	// "og ni dlrow olleh"
	// "og" -> "go"
	// ...
	// "olleh" -> "hello"
	// result: "go in world hello"
	i := 0
	for i < len(words) {
		if words[i] == ' ' {
			i++
			continue
		}

		wordEnd := i
		for wordEnd+1 < len(words) && words[wordEnd+1] != ' ' {
			wordEnd++
		}

		r := wordEnd
		for i < r {
			words[i], words[r] = words[r], words[i]
			i++
			r--
		}
		i = wordEnd + 1
	}
}

