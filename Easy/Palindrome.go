package Easy

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	y := x
	rev := 0
	for x != 0 {
		rev = (rev * 10) + (x % 10)
		x /= 10

	}
	return rev == y
}
