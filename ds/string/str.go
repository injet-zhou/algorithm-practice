package string

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		return s[0:1]
	}
	data := []rune(s)
	l,r :=0,0
	if len(data)%2 == 0{
		l = len(data)/2 -1
		r = l+1
	} else {
		l = len(data)/2
		r = l
	}
	for l >0 && r<= len(data)-1  && data[l] == data[r] {
		l--
		r++
	}
	if r == len(data) - 1 && data[l] == data[r] {
		r++
	}
	if l == 0 && r != len(data) &&  data[l] != data[r] {
		l++
	}
	return s[l:r]
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	begin := false
	negative := false
	result := 0
	for i, str := range s {
		if i == 0 {
			if (str < 48 || str > 57) && str != 32 && str != 43 && str != 45 {
				return 0
			}
			if str == 45 {
				negative = true
				begin = true
			}
			if str == 43 {
				begin = true
			}
			if str >= 48 && str <= 57 {
				result = int(str - 48)
				begin = true
				continue
			}
		}
		if str == 45 {
			negative = true
			continue
		}
		if !begin && str == 32{
			continue
		}
		if begin {
			if str < 48 || str > 57 {
				break
			}
		}
		if str >= 48 && str <= 57 {
			begin = true
			result *= 10
			result += int(str - 48)
		}
	}
	if negative {
		result =  -result
	}
	if result <= -2147483648 {
		return -2147483648
	}
	if result >= 2147483647 {
		return 2147483647
	}
	return result
}
