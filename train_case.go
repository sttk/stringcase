// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to train case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is train case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are eliminated as word separators.
func TrainCase(input string) string {
	result := make([]rune, 0, len(input)*2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, toAsciiUpperCase(ch))
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					prev = toAsciiUpperCase(prev)
				}
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', toAsciiUpperCase(ch))
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) {
			switch flag {
			case ChIsNextOfSepMark:
				result = append(result, '-', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsNextOfKeepedMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		}
	}

	return string(result)
}

// Converts a string to train case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is train case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are replaced to hyphens.
func TrainCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input)*2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if strings.ContainsRune(seps, ch) {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		} else if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, toAsciiUpperCase(ch))
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					prev = toAsciiUpperCase(prev)
				}
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', toAsciiUpperCase(ch))
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else {
			if flag == ChIsNextOfSepMark {
				result = append(result, '-', ch)
			} else {
				result = append(result, ch)
			}
			flag = ChIsNextOfKeepedMark
		}
	}

	return string(result)
}

// Converts a string to train case using characters other than the specified
// characters as separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is train case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters other than the specified characters as
// the second argument of this function are regarded as word separators and
// are replaced to hyphens.
func TrainCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input)*2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, ch)
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsFirstOfStr:
				result = append(result, toAsciiUpperCase(ch))
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					prev = toAsciiUpperCase(prev)
				}
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', toAsciiUpperCase(ch))
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '-', ch)
			} else {
				result = append(result, ch)
			}
			flag = ChIsNextOfKeepedMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		}
	}

	return string(result)
}
