func reverseWords(s string) string {

	words := [][]byte{}

	i := 0
	for i < len(s) {
		word := []byte{}
		if s[i] != ' ' {
			for i < len(s) && s[i] != ' ' {
				word = append(word, s[i])
				i++
			}

			words = append(words, word)
		}
		i++
	}

	if len(words) == 0 {
		return ""
	}

	reverse(words)
	return buildSentence(words)

}

func reverse(a [][]byte) {
	j := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		a[i], a[j-i] = a[j-i], a[i]
	}

}

func buildSentence(a [][]byte) string {
	ans := []byte{}
	for i, word := range a {
		ans = append(ans, word...)
		if i != len(a)-1 {
			ans = append(ans, ' ')
		}
	}

	return string(ans)

}