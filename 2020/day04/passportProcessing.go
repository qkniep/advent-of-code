package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fields = make(map[string]string)
	var valuesPresent = 0
	var valuesValid = 0

	// read and validate the passport inputs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if checkPresent(fields) {
				valuesPresent++
				valuesValid += checkValid(fields)
			}
			fields = make(map[string]string)
		} else {
			for _, entry := range strings.Fields(line) {
				kv := strings.Split(entry, ":")
				fields[kv[0]] = kv[1]
			}
		}
	}
	if checkPresent(fields) {
		valuesPresent++
		valuesValid += checkValid(fields)
	}

	fmt.Printf("All required fields: %d\n", valuesPresent)
	fmt.Printf("+ all entries valid: %d\n", valuesValid)
}

// Returns `true` iff all required fields are present.
func checkPresent(fields map[string]string) bool {
	var required = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, r := range required {
		if fields[r] == "" {
			return false
		}
	}
	return true
}

// Returns 1 if all fields have valid values according to their specific rules, 0 otherwise.
func checkValid(fields map[string]string) int {
	validHeight := false
	var height = 0
	_, err := fmt.Sscanf(fields["hgt"], "%din", &height)
	if err == nil && height >= 59 && height <= 76 {
		validHeight = true
	}
	_, err = fmt.Sscanf(fields["hgt"], "%dcm", &height)
	if err == nil && height >= 150 && height <= 193 {
		validHeight = true
	}
	if !validHeight {
		return 0
	}
	if !checkValidNum(fields["byr"], 4, 1920, 2002, 10) ||
		!checkValidNum(fields["iyr"], 4, 2010, 2020, 10) ||
		!checkValidNum(fields["eyr"], 4, 2020, 2030, 10) ||
		fields["hcl"][0] != '#' ||
		!checkValidNum(fields["hcl"][1:], 6, 0, 16_777_216, 16) ||
		!checkValidNum(fields["pid"], 9, 0, 999_999_999, 10) {
		return 0
	}
	switch fields["ecl"] {
	case
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		return 0
	}
	return 1
}

// Checks whether `s` is exactly `digits` long and, when interpreted as a number in base `base`, is
// between `min` and `max`. Returns `true` iff this is the case.
func checkValidNum(s string, digits int, min int64, max int64, base int) bool {
	if len(s) != digits {
		return false
	}
	num, err := strconv.ParseInt(s, base, 32)
	return err == nil && num >= min && num <= max
}
