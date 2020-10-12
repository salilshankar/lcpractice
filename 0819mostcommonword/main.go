func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	words := strings.FieldsFunc(paragraph, f)

	m := make(map[string]int)
	for i, bannedWord := range banned {
		m[bannedWord] = i
	}

	ans := ""
	ansCount := 0

	m1 := make(map[string]int)

	for _, word := range words {
		if _, banned := m[word]; !banned {
			if _, found := m1[word]; found {
				m1[word] = m1[word] + 1
			} else {
				m1[word] = 1
			}

			v := m1[word]
			if v > ansCount {
				ansCount = v
				ans = word
			}
		}
	}

	return ans

}