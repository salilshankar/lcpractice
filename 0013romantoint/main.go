import "strings"

func romanToInt(s string) int {

	type pair struct {
		p string
		v int
	}

	roman := []pair{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},

		{"I", 1},
	}

	result := 0
	for _, p := range roman {
		for strings.HasPrefix(s, p.p) {
			result = result + p.v
			s = s[len(p.p):]

		}
	}

	return result
}