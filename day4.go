package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//byr (Birth Year)
//iyr (Issue Year)
//eyr (Expiration Year)
//hgt (Height)
//hcl (Hair Color)
//ecl (Eye Color)
//pid (Passport ID)
//cid (Country ID)

func init_passport() map[string]string {
	return map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		//"cid": "",
	}
}

var passport_regex = map[string]func(string) bool{
	"byr": func(value string) bool {
		matched, _ := regexp.Match("^[0-9]{4}$", []byte(value))
		if matched {
			valuei, _ := strconv.Atoi(value)
			if valuei >= 1920 && valuei <= 2002 {
				return true
			}
		}
		return false
	},
	"iyr": func(value string) bool {
		matched, _ := regexp.Match("^[0-9]{4}$", []byte(value))
		if matched {
			valuei, _ := strconv.Atoi(value)
			if valuei >= 2010 && valuei <= 2020 {
				return true
			}
		}
		return false
	},
	"eyr": func(value string) bool {
		matched, _ := regexp.Match("^[0-9]{4}$", []byte(value))
		if matched {
			valuei, _ := strconv.Atoi(value)
			if valuei >= 2020 && valuei <= 2030 {
				return true
			}
		}
		return false
	},
	"hgt": func(value string) bool {
		matchedcm, _ := regexp.Match("^[0-9]{3}cm$", []byte(value))
		matchedin, _ := regexp.Match("^[0-9]{2}in$", []byte(value))
		if matchedcm {
			valuei, _ := strconv.Atoi(value[:3])
			if valuei >= 150 && valuei <= 193 {
				return true
			}
		} else if matchedin {
			valuei, _ := strconv.Atoi(value[:2])
			if valuei >= 59 && valuei <= 76 {
				return true
			}
		}
		return false
	},
	"hcl": func(value string) bool {
		matched, _ := regexp.Match("^#[0-9a-f]{6}$", []byte(value))
		return matched
	},
	"ecl": func(value string) bool {
		matched, _ := regexp.Match("^(amb|blu|brn|gry|grn|hzl|oth)$", []byte(value))
		return matched
	},
	"pid": func(value string) bool {
		matched, _ := regexp.Match("^[0-9]{9}$", []byte(value))
		return matched
	},
}

func passport_is_valid(passport map[string]string, strict bool) bool {
	// first validate only if all keys are present
	for _, value := range passport {
		if value == "" { //&& key != "cid" {
			return false
		}
	}
	if !strict {
		return true
	}
	// also validate the values
	for key, value := range passport {
		if !passport_regex[key](value) {
			return false
		}
	}

	return true
}

func main() {
	file, _ := os.Open("day4_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	valid_nonstrict := 0
	nonvalid_nonstrict := 0
	valid_strict := 0
	nonvalid_strict := 0
	num_line := 0
	passport := init_passport()
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			// new passport
			// check if last one is valid now that it is parsed fully
			if passport_is_valid(passport, false) {
				valid_nonstrict += 1
			} else {
				nonvalid_nonstrict += 1
			}
			if passport_is_valid(passport, true) {
				valid_strict += 1
			} else {
				nonvalid_strict += 1
			}
			// clear the passport object
			passport = init_passport()
		} else {
			// split on spaces
			for _, prop := range strings.Fields(line) {
				// split on seperator
				parts := strings.Split(prop, ":")
				if parts[0] != "cid" {
					passport[parts[0]] = parts[1]
				}
			}
		}
		num_line += 1
	}
	// due to the nature of the scanner the last passport check is not triggered, lazy solution is to copy that here
	if passport_is_valid(passport, false) {
		valid_nonstrict += 1
	} else {
		nonvalid_nonstrict += 1
	}
	if passport_is_valid(passport, true) {
		valid_strict += 1
	} else {
		nonvalid_strict += 1
	}
	fmt.Println("Valid/nonvalid passports non strict: ", valid_nonstrict, nonvalid_nonstrict)
	fmt.Println("Valid/nonvalid passports strict: ", valid_strict, nonvalid_strict)

}
