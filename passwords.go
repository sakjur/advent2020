package advent2020

import (
	"fmt"
	"regexp"
	"strconv"
)

var passwordLineRegexp = regexp.MustCompile("(\\d+)-(\\d+) ([a-zA-Z]): (\\w+)")

type PasswordRule struct {
	Low  int
	High int
	Char rune
}

func (r PasswordRule) Validate(password string) bool {
	count := 0
	for _, c := range password {
		if c == r.Char {
			count++
		}
	}

	return count >= r.Low && count <= r.High
}

func (r PasswordRule) ValidateIndexes(password string) bool {
	iH := r.High - 1
	iL := r.Low - 1
	if iH >= len(password) || iL >= len(password) {
		return false
	}

	low := rune(password[iL]) == r.Char
	high := rune(password[iH]) == r.Char
	return (high && !low) || (low && !high)
}

func ParsePasswordRule(line string) (PasswordRule, string, error) {
	submatches := passwordLineRegexp.FindStringSubmatch(line)

	if len(submatches) != 5 {
		return PasswordRule{}, "", fmt.Errorf("expected exactly 5 fields on match, got %d", len(submatches))
	}
	min, err := strconv.Atoi(submatches[1])
	if err != nil {
		return PasswordRule{}, "", err
	}

	max, err := strconv.Atoi(submatches[2])
	if err != nil {
		return PasswordRule{}, "", err
	}

	char := rune(submatches[3][0])

	return PasswordRule{
			Low:  min,
			High: max,
			Char: char,
		},
		submatches[4],
		nil
}

func ValidateLength(lines <-chan string) (length int, indexes int) {
	for line := range lines {
		rule, pwd, err := ParsePasswordRule(line)
		if err != nil {
			panic(err)
		}

		if rule.Validate(pwd) {
			length++
		}
		if rule.ValidateIndexes(pwd) {
			indexes++
		}
	}

	return length, indexes
}
