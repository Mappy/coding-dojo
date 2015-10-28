package main

func Roman(number int) string {
	var result=""
	replace(&number, &result, 1000, "M")
	replace(&number, &result, 900, "CM")
	replace(&number, &result, 500, "D")
	replace(&number, &result, 400, "CD")
	replace(&number, &result, 100, "C")
	replace(&number, &result, 90, "XC")
	replace(&number, &result, 50, "L")
	replace(&number, &result, 40, "XL")
	replace(&number, &result, 10, "X")
	replace(&number, &result, 9, "IX")
	replace(&number, &result, 5, "V")
	replace(&number, &result, 4, "IV")
	replace(&number, &result, 1, "I")
	return result
}

func replace(number *int, result *string, inc int, romanChar string) {
	for (*number >= inc) {
		*result += romanChar
		*number = *number -inc
	}
}
