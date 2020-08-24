import "strings"

func reverseWords(s string) string {
	strarr := strings.Fields(s)
	strarr = reverse(strarr)

	return strings.Join(strarr, " ")

}

func reverse(a []string) []string {
	j := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		a[i], a[j-i] = a[j-i], a[i]
	}

	return a
}