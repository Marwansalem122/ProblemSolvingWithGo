package Easy

import "strings"

func lengthOfLastWord(s string) int {
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])

}
