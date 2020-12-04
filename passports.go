package advent2020

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type TravelDocument map[string]string

func (t TravelDocument) MissingFields() error {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	missing := []string{}
	for _, field := range requiredFields {
		if _, exists := t[field]; !exists {
			missing = append(missing, field)
		}
	}
	if len(missing) != 0 {
		return fmt.Errorf("missing fields: %v", missing)
	}
	return nil
}

var hclRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
var eclRegexp = regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
var pidRegexp = regexp.MustCompile("^[0-9]{9}$")

func (t TravelDocument) Valid() bool {
	if err := t.MissingFields(); err != nil {
		return false
	}

	errs := []error{}
	errs = append(errs, t.NumericRule("byr", 1920, 2002))
	errs = append(errs, t.NumericRule("iyr", 2010, 2020))
	errs = append(errs, t.NumericRule("eyr", 2020, 2030))

	if strings.HasSuffix(t["hgt"], "in") {
		t["hgt"] = strings.TrimSuffix(t["hgt"], "in")
		errs = append(errs, t.NumericRule("hgt", 59, 76))
	} else if strings.HasSuffix(t["hgt"], "cm") {
		t["hgt"] = strings.TrimSuffix(t["hgt"], "cm")
		errs = append(errs, t.NumericRule("hgt", 150, 193))
	} else {
		errs = append(errs, fmt.Errorf("[hgt] must be suffixed by either 'in' or 'cm': %v", t["hgt"]))
	}

	if !hclRegexp.MatchString(t["hcl"]) {
		errs = append(errs, fmt.Errorf("[hcl] must match regexp '^#[0-9a-f]{6}$'"))
	}

	if !eclRegexp.MatchString(t["ecl"]) {
		errs = append(errs, fmt.Errorf("[ecl] must be either of amb|blu|brn|gry|grn|hzl|oth"))
	}

	if !pidRegexp.MatchString(t["pid"]) {
		errs = append(errs, fmt.Errorf("[pid] must be a nine digit number including leading zeroes"))
	}

	for _, err := range errs {
		if err != nil {
			return false
		}
	}
	return true
}

func (t TravelDocument) NumericRule(field string, min, max int) error {
	value, err := strconv.Atoi(t[field])
	if err != nil {
		return fmt.Errorf("[%s] expected '%s' to be a numeric value", field, t[field])
	}
	if value < min || value > max {
		return fmt.Errorf("[%s] expected '%s' to be between %d and %d", field, t[field], min, max)
	}
	return nil
}

func (t TravelDocument) Empty() bool {
	return len(t) == 0
}

func ScanDocuments(r io.Reader) []TravelDocument {
	linescan := bufio.NewScanner(r)

	documents := []TravelDocument{}
	current := TravelDocument{}

	for linescan.Scan() {
		line := linescan.Text()
		if line == "" {
			documents = append(documents, current)
			current = TravelDocument{}
			continue
		}

		columns := strings.Split(line, " ")
		for _, column := range columns {
			kv := strings.Split(column, ":")
			if len(kv) != 2 {
				panic(fmt.Sprintln("expected 2 components in ", kv))
			}
			current[kv[0]] = kv[1]
		}
	}

	if !current.Empty() {
		documents = append(documents, current)
	}
	return documents
}
