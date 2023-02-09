package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	var resultString string
	fmt.Println("Character\tOrdinal Value\tOrdinal + 7\tResulting Character")
	for _, char := range input {
		charValue := int(char)
		resultValue := charValue - 7
		if resultValue > 122 {
			resultValue -= 26
		} else if resultValue < 97 {
			resultValue += 26
		}
		resultChar := string(rune(resultValue))
		resultString += resultChar

		fmt.Println(string(char), "\t\t", charValue, "\t\t", resultValue, "\t\t", resultChar)
	}
	fmt.Println("Resulting string:", resultString)
}
