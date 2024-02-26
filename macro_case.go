// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Converts a string to macro case.
//
// This function takes a string as its argument, then returns a string of which
// the case style is macro case.
//
// This function targets the upper and lower cases of ASCII alphabets for
// capitalization, and all characters except ASCII alphabets and ASCII numbers
// are eliminated as word separators.
func MacroCase(input string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOthers
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
				result = append(result, '_', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, toAsciiUpperCase(ch))
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', toAsciiUpperCase(ch))
			default:
				result = append(result, toAsciiUpperCase(ch))
			}
			flag = ChIsOthers
		} else if isAsciiDigit(ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '_', ch)
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

// Converts a string to macro case using the specified characters as
// separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is macro case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters specified as the second argument of this
// function are regarded as word separators and are replaced to underscores.
func MacroCaseWithSep(input, seps string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOthers
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
				result = append(result, '_', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, toAsciiUpperCase(ch))
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', toAsciiUpperCase(ch))
			default:
				result = append(result, toAsciiUpperCase(ch))
			}
			flag = ChIsOthers
		} else {
			if flag == ChIsNextOfSepMark {
				result = append(result, '_', ch)
			} else {
				result = append(result, ch)
			}
			flag = ChIsNextOfKeepedMark
		}
	}

	return string(result)
}

// Converts a string to macro case using characters other than the specified
// characters as separators.
//
// This function takes a string as its argument, then returns a string of which
// the case style is macro case.
//
// This function targets only the upper and lower cases of ASCII alphabets for
// capitalization, and the characters other than the specified characters as
// the second argument of this function are regarded as word separators and
// are replaced to underscores.
func MacroCaseWithKeep(input, keeped string) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeepedMark
		ChIsOthers
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
				result = append(result, '_', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, toAsciiUpperCase(ch))
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', toAsciiUpperCase(ch))
			default:
				result = append(result, toAsciiUpperCase(ch))
			}
			flag = ChIsOthers
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '_', ch)
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
