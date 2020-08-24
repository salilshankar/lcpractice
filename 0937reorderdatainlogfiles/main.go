package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arrInput := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	fmt.Println(arrInput)
	fmt.Println(reorderLogFiles(arrInput))
}

func reorderLogFiles(logs []string) []string {
	numLogs := []string{}
	for i, v := range logs {
		temp := strings.Fields(v)

		if temp[1][0] <= '9' {
			numLogs = append(numLogs, v)
		}
		temp = append(temp[1:], temp[0])
		reorderedValue := strings.Join(temp, " ")
		logs[i] = reorderedValue

	}

	sort.Strings(logs)

	for i, v := range logs {
		temp := strings.Fields(v)
		tempArr := []string{temp[len(temp)-1]}
		temp = append(tempArr, temp[:len(temp)-1]...)
		reorderedValue := strings.Join(temp, " ")
		logs[i] = reorderedValue
	}

	if len(numLogs) != 0 {
		logs = append(logs, numLogs...)
		logs = logs[len(numLogs):]
	}

	return logs

}
