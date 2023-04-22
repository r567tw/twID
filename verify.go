package twID // Package twID provides a function to verify Taiwanese ID numbers.

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var local_table = map[string]string{
	"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15", "G": "16", "H": "17", "I": "34",
	"J": "18", "K": "19", "L": "20", "M": "21", "N": "22", "O": "35", "P": "23", "Q": "24", "R": "25",
	"S": "26", "T": "27", "U": "28", "V": "29", "W": "32", "X": "30", "Y": "31", "Z": "33",
}

// Package twID provides a function to verify Taiwanese ID numbers.
// Verify checks if a Taiwanese ID number is valid or not.
// Returns true if the ID number is valid, otherwise false.
func Verify(IDNumber string) bool {
	id := strings.ToUpper(IDNumber)

	match, err := regexp.MatchString("^[A-Z][1289][0-9]{8}$", id)

	if err != nil {
		fmt.Println(err)
		panic("Verify ERROR")
	}

	if !match {
		return false
	}

	// Convert the first letter of the ID number to its corresponding number.
	numbers := strings.Split(id, "")
	new_id := strings.Replace(id, numbers[0], local_table[numbers[0]], 1)

	numbers = strings.Split(new_id, "")
	intNumbers := []int{}

	// Convert the ID number to an array of integers.
	for _, str := range numbers {
		i, _ := strconv.Atoi(str)
		intNumbers = append(intNumbers, i)
	}

	// Check if the checksum is correct.
	result := 0
	result += intNumbers[0]

	for x := 0; x < 9; x++ {
		result += intNumbers[x+1] * (9 - x)
	}

	result += intNumbers[10]

	return result%10 == 0
}
