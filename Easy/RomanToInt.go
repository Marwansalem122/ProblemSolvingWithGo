package Easy

func romanToInt(s string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && values[rune(s[i])] < values[rune(s[i+1])] {
			result -= values[rune(s[i])]
		} else {
			result += values[rune(s[i])]
		}
	}

	return result
}

/*

func romanToInt(s string) int {
	var result int = 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == 'I' && s[i+1] == 'V' {
			result += 4
			i += 1
		} else if i < len(s)-1 && s[i] == 'I' && s[i+1] == 'X' {
			result += 9
			i += 1
		} else if i < len(s)-1 && s[i] == 'X' && s[i+1] == 'L' {
			result += 40
			i += 1
		} else if i < len(s)-1 && s[i] == 'X' && s[i+1] == 'C' {
			result += 90
			i += 1
		} else if i < len(s)-1 && s[i] == 'C' && s[i+1] == 'D' {
			result += 400
			i += 1
		} else if i < len(s)-1 && s[i] == 'C' && s[i+1] == 'M' {
			result += 900
			i += 1
		} else {
			result += value(rune(s[i]))
		}
	}
	return result
}

func value(s rune) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000

	}
	return 0
}

*/
