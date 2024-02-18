// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to snake case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is snake case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are eliminated as word separators.
func SnakeCase(input string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: one and two chars before are upper
	// 3: previous char is mark
	// 4: other

	for _, r := range input {
		if isAsciiUpperCase(r) {
			switch flag {
			case 0:
				flag = 1
			case 1, 2:
				flag = 2
			default:
				flag = 1
				result = append(result, '_')
			}
			result = append(result, toAsciiLowerCase(r))
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 2:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev)
			case 3:
				result = append(result, '_')
			}
			flag = 4
			result = append(result, r)
		} else if isAsciiDigit(r) {
			if flag == 3 {
				result = append(result, '_')
			}
			flag = 4
			result = append(result, r)
		} else {
			if flag != 0 {
				flag = 3
			}
		}
	}

	return string(result)
}

// Converts a string to snake case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is snake case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are replaced to underscores.
func SnakeCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: one and two chars before are upper
	// 3: previous char is mark (separator)
	// 4: previous char is mark (keeped)
	// 5: other

	for _, r := range input {
		if strings.ContainsRune(seps, r) {
			if flag != 0 {
				flag = 3
			}
		} else if isAsciiUpperCase(r) {
			switch flag {
			case 0:
				flag = 1
			case 1, 2:
				flag = 2
			default:
				flag = 1
				result = append(result, '_')
			}
			result = append(result, toAsciiLowerCase(r))
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 2:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev)
			case 3, 4:
				result = append(result, '_')
			}
			flag = 5
			result = append(result, r)
		} else if isAsciiDigit(r) {
			switch flag {
			case 3, 4:
				result = append(result, '_')
			}
			flag = 5
			result = append(result, r)
		} else {
			if flag == 3 {
				result = append(result, '_')
			}
			flag = 4
			result = append(result, r)
		}
	}

	return string(result)
}

func SnakeCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: one and two chars before are upper
	// 3: previous char is mark (separator)
	// 4: previous char is mark (keeped)
	// 5: other

	for _, r := range input {
		if isAsciiUpperCase(r) {
			switch flag {
			case 0:
				flag = 1
			case 1, 2:
				flag = 2
			default:
				flag = 1
				result = append(result, '_')
			}
			result = append(result, toAsciiLowerCase(r))
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 2:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev)
			case 3, 4:
				result = append(result, '_')
			}
			flag = 5
			result = append(result, r)
		} else if isAsciiDigit(r) {
			switch flag {
			case 3, 4:
				result = append(result, '_')
			}
			flag = 5
			result = append(result, r)
		} else if strings.ContainsRune(keeped, r) {
			if flag == 3 {
				result = append(result, '_')
			}
			flag = 4
			result = append(result, r)
		} else {
			if flag != 0 {
				flag = 3
			}
		}
	}

	return string(result)
}
