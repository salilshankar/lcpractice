func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[rune]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"

	ans := []string{""}

	for _, digit := range digits {
		ans = expand(ans, m[digit])
	}

	return ans

}

func expand(anslist []string, node string) []string {
	next := []string{}
	for _, ans := range anslist {
		for _, char := range node {
			next = append(next, ans+string(char))
		}
	}

	return next
}