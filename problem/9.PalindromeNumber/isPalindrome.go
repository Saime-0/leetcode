package main

func isPalindrome(x int) bool {
	// Мгновенно вернуть результат
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	// 1. Определить количество разрядов
	length := 1
	divider := 10
	for {
		if x/divider == 0 {
			break
		}
		length++
		divider *= 10
	}

	// 2. разряды поместить в слайс
	numbers := make([]int, length)
	remains := x
	for i := 0; i < length; i++ {
		numbers[i] = remains % 10
		remains /= 10
	}

	// 3. Линейно определить палиндром
	firstRightMirrored := length / 2
	rightOffset := length % 2 // Количество разрядов нечетное
	for i := 0; firstRightMirrored+rightOffset+i < length; i++ {
		right := numbers[i+firstRightMirrored+rightOffset]
		left := numbers[firstRightMirrored-i-1]
		if left != right {
			return false
		}
	}

	return true
}
