// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to kebab case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is kebab case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are eliminated as word separators.
func KebabCase(input string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-')
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) {
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

// Converts a string to kebab case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is kebab case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are replaced to hyphens.
func KebabCaseWithSep(input, seps string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) || !strings.ContainsRune(seps, ch) {
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

// Converts a string to kebab case using characters other than the specified
// characters as separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is kebab case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters other than the specified characters as
// the second argument of this function are regarded as word separators and
// are replaced to hyphens.
func KebabCaseWithKeep(input, keeped string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '-', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '-', ch)
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
