// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to pascal case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is pascal case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are eliminated as word separators.
func PascalCase(input string) string {
	result := make([]rune, 0, len(input))

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: previous char is mark
	// 3: other

	for _, r := range input {
		if isAsciiUpperCase(r) {
			switch flag {
			case 1:
				flag = 1
				result = append(result, toAsciiLowerCase(r))
			default:
				flag = 1
				result = append(result, r)
			}
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 1:
				flag = 3
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, r)
			case 0, 2:
				flag = 1
				result = append(result, toAsciiUpperCase(r))
			default:
				flag = 3
				result = append(result, r)
			}
		} else if isAsciiDigit(r) {
			switch flag {
			case 0, 2:
				flag = 1
				result = append(result, r)
			default:
				flag = 3
				result = append(result, r)
			}
		} else {
			if flag != 0 {
				flag = 2
			}
		}
	}

	return string(result)
}

// Converts a string to pascal case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is pascal case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are eliminated.
func PascalCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input))

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: previous char is mark
	// 3: other

	for _, r := range input {
		if strings.ContainsRune(seps, r) {
			if flag != 0 {
				flag = 2
			}
		} else if isAsciiUpperCase(r) {
			switch flag {
			case 1:
				flag = 1
				result = append(result, toAsciiLowerCase(r))
			default:
				flag = 1
				result = append(result, r)
			}
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 1:
				flag = 3
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, r)
			case 0, 2:
				flag = 1
				result = append(result, toAsciiUpperCase(r))
			default:
				flag = 3
				result = append(result, r)
			}
		} else if isAsciiDigit(r) {
			switch flag {
			case 0, 2:
				flag = 1
				result = append(result, r)
			default:
				flag = 3
				result = append(result, r)
			}
		} else {
			flag = 2
			result = append(result, r)
		}
	}

	return string(result)
}

// Converts a string to pascal case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is pascal case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are eliminated.
func PascalCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input))

	var flag uint8 = 0
	// 0: first char
	// 1: previous char is upper
	// 2: previous char is mark
	// 3: other

	for _, r := range input {
		if isAsciiUpperCase(r) {
			switch flag {
			case 1:
				flag = 1
				result = append(result, toAsciiLowerCase(r))
			default:
				flag = 1
				result = append(result, r)
			}
		} else if isAsciiLowerCase(r) {
			switch flag {
			case 1:
				flag = 3
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, r)
			case 0, 2:
				flag = 1
				result = append(result, toAsciiUpperCase(r))
			default:
				flag = 3
				result = append(result, r)
			}
		} else if isAsciiDigit(r) {
			switch flag {
			case 0, 2:
				flag = 1
				result = append(result, r)
			default:
				flag = 3
				result = append(result, r)
			}
		} else if strings.ContainsRune(keeped, r) {
			flag = 2
			result = append(result, r)
		} else {
			if flag != 0 {
				flag = 2
			}
		}
	}

	return string(result)
}
