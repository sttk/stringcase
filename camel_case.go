// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to camel case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is camel case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASII alphabets and ASII numbers
// are eliminated as word separators.
func CamelCase(input string) string {
	result := make([]rune, 0, len(input))

	const (
		ChIsFirstOfStr = iota
		ChIsInFirstWord
		ChIsNextOfUpper
		ChIsNextOfMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsInFirstWord:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsInFirstWord
			case ChIsNextOfUpper:
				result = append(result, toAsciiLowerCase(ch))
				//flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, ch)
				flag = ChIsOther
			case ChIsNextOfMark:
				result = append(result, toAsciiUpperCase(ch))
				flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsOther
			}
		} else if isAsciiDigit(ch) {
			result = append(result, ch)
			flag = ChIsNextOfMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfMark
			}
		}
	}

	return string(result)
}

// Converts a string to camel case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is camel case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are eliminated.
func CamelCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input))

	const (
		ChIsFirstOfStr = iota
		ChIsInFirstWord
		ChIsNextOfUpper
		ChIsNextOfMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsInFirstWord:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsInFirstWord
			case ChIsNextOfUpper:
				result = append(result, toAsciiLowerCase(ch))
				//flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, ch)
				flag = ChIsOther
			case ChIsNextOfMark:
				result = append(result, toAsciiUpperCase(ch))
				flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsOther
			}
		} else if isAsciiDigit(ch) || !strings.ContainsRune(seps, ch) {
			result = append(result, ch)
			flag = ChIsNextOfMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfMark
			}
		}
	}

	return string(result)
}

// Converts a string to camel case using characters other than the specified
// characters as separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is camel case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters other than the specified characters as
// the second argument of this function are regarded as word separators and
// are eliminated.
func CamelCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input))

	const (
		ChIsFirstOfStr = iota
		ChIsInFirstWord
		ChIsNextOfUpper
		ChIsNextOfMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsInFirstWord:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsInFirstWord
			case ChIsNextOfUpper:
				result = append(result, toAsciiLowerCase(ch))
				//flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					result[n-1] = toAsciiUpperCase(prev)
				}
				result = append(result, ch)
				flag = ChIsOther
			case ChIsNextOfMark:
				result = append(result, toAsciiUpperCase(ch))
				flag = ChIsNextOfUpper
			default:
				result = append(result, ch)
				flag = ChIsOther
			}
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			result = append(result, ch)
			flag = ChIsNextOfMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfMark
			}
		}
	}

	return string(result)
}
