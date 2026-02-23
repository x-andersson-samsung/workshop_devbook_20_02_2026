package exercise2

//Write a function called AtoI accepting a string representation of a number and returning an int
//You can assume that the passed value will be correct.
func AtoI(input string) int {
	isNegative := false
	if input[0] == '-' {
		isNegative = true
		input = input[1:]
	}

	var sum int
	for _, char := range input {
		nextDigit := int(char - '0')
		sum = 10*sum + nextDigit
	}

	if isNegative {
		sum = -sum
	}

	return sum
}
