import "strings"

func intToRoman(num int) string {

	thous := []string{"", "M", "MM", "MMM"}
	hunds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	result := []string{thous[num/1000], hunds[(num%1000)/100], tens[(num%100)/10], ones[num%10]}

	return strings.Join(result, "")

}