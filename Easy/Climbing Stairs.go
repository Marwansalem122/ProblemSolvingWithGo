package Easy

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
