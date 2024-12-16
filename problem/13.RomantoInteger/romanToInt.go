package main

func romanToInt(s string) int {
	rtoi := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	prevRoman := s[len(s)-1]
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		integer := rtoi[r]
		prevInteger := rtoi[prevRoman]
		if integer < prevInteger {
			result -= integer
		} else {
			result += integer
			prevRoman = r
		}
	}
	return result
}
