package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

/*
Task: Given a list of 4 integers, find all possible valid 24 hour times (eg: 12:34)
that the given integers can be assembled into and return the total number of valid times.
You can not use the same number twice.
Times such as 34:12 and 12:60 are not valid.
Provided integers can not be negative.
Notes: Input integers can not be negative.
Input integers can yeald 0 possible valid combinations.
Example:
	Input: [1, 2, 3, 4]
	Valid times: ["12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "23:14", "23:41", "21:34", "21:43"]
	Return: 10
*/

// returns all combinations possible regarless time constraint
func allCombinations(d []int) []string {

	//  store all combinations possible
	// 4 digits so 4 exponent 4 possibilities (256)
	acp := make([]string, 0, 256)
	for i := range d {
		for j := range d {
			for k := range d {
				for l := range d {
					// digit position
					dp1 := strconv.Itoa(d[i])
					dp2 := strconv.Itoa(d[j])
					dp3 := strconv.Itoa(d[k])
					dp4 := strconv.Itoa(d[l])
					oc := dp1 + dp2 + dp3 + dp4 // one combination
					acp = append(acp, oc)       // all combinations possible
				}
			}
		}
	}

	return acp
}

// returns all valid times from all combinations
func validTimes(d []string, digits []int) (combinations []string) {

	// Loop over all combinations to remove non valid time
	for _, e := range d {

		d1 := e[0:1] // first digit position
		d2 := e[1:2] // second digit position
		d3 := e[2:3] // third digit position

		// the first digit of a 24 hour time can only be [0 1 2]
		if d1 < "3" {
			// case first digit is 2
			if d1 == "2" {
				// the second digit of a 24 hour time can only be [0 1 2 3] when first digit is 2
				if d2 < "4" {
					if isValid := validTime(e, digits); isValid {
						combinations = append(combinations, e)
					}
				}
			} else {
				// the third digit of a 24 hour time can only be [0 1 2 3 4 5]
				if d3 < "6" {
					if isValid := validTime(e, digits); isValid {
						combinations = append(combinations, e)
					}
				}
			}
		}
	}

	return
}

// returns wether a combination contains all digits
func validTime(e string, digits []int) (isValid bool) {
	// store if all digits are in the element
	// we don't want the same digit to be used twice
	valid := make([]bool, 0, 4)

	// loop over all digits to check if it is contains in the element
	for _, d := range digits {
		if strings.Contains(e, strconv.Itoa(d)) {
			valid = append(valid, true)
		} else {
			valid = append(valid, false)
		}
	}

	// if valid contains one false it is not valid
	isNotValid := slices.Contains(valid, false) // isNotValid = false => isValid = true

	return !isNotValid
}

func possibleTimes(digits []int) (x int) {
	vt := validTimes(allCombinations(digits), digits) // valid times
	x = len(vt)                                       // total number of valid times

	return
}

func main() {
	// Example test cases.
	fmt.Println(possibleTimes([]int{1, 2, 3, 4}))
	fmt.Println(possibleTimes([]int{9, 1, 2, 0}))
	fmt.Println(possibleTimes([]int{2, 2, 1, 9}))

}
