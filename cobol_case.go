// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to cobol case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is cobol case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are replaced to hyphens as word separators.
func CobolCase(input string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

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
				result = append(result, ch)
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-')
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-')
			}
			result = append(result, toAsciiUpperCase(ch))
			flag = ChIsOther
		} else if isAsciiDigit(ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '-')
			}
			result = append(result, ch)
			flag = ChIsNextOfKeepedMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		}
	}

	return string(result)
}

// Converts a string to cobol case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is cobol case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are replaced to hyphens.
func CobolCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

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
				result = append(result, ch)
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, toAsciiUpperCase(ch))
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', toAsciiUpperCase(ch))
			default:
				result = append(result, toAsciiUpperCase(ch))
			}
			flag = ChIsOther
		} else {
			switch flag {
			case ChIsNextOfSepMark:
				result = append(result, '-', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsNextOfKeepedMark
		}
	}

	return string(result)
}

// Converts a string to cobol case using characters other than the specified
// characters as separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is cobol case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters other than the specified characters as
// the second argument of this function are regarded as word separators and
// are replaced to hyphens.
func CobolCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

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
				result = append(result, ch)
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-')
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, toAsciiUpperCase(ch))
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', toAsciiUpperCase(ch))
			default:
				result = append(result, toAsciiUpperCase(ch))
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '-')
			}
			result = append(result, ch)
			flag = ChIsNextOfKeepedMark
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		}
	}

	return string(result)
}
