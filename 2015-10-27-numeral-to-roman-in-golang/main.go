package main

func Aux(number int, result string) string {
	if number == 0 {
		return result
	}
	if number >= 500 {
		return Aux(number - 500, result + "D")
	}
	if number >= 400 {
		return Aux(number - 400, result + "CD")
	}
	if number >= 100 {
		return Aux(number - 100, result + "C")
	}
	if number >= 90 {
		return Aux(number - 90, result + "XC")
	}
	if number >= 50 {
		return Aux(number - 50, result + "L")
	}
	if number >= 10 {
		return Aux(number - 10, result + "X")
	}
	if number >= 9 {
		return Aux(number - 9, result + "IX")
	}
	if number >= 5 {
		return Aux(number - 5, result + "V")
	}
	if number >= 4 {
		return Aux(number - 4, result + "IV")
	}
	return Aux(number - 1, result + "I")
}

func Roman(number int) string {
	return Aux(number, "")
}
