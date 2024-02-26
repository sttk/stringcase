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
				result = append(result, '_', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '_', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
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

func SnakeCaseWithKeep(input, keeped string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper:
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			default:
				result = append(result, '_', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			switch flag {
			case ChIsNextOfContdUpper:
				n := len(result)
				prev := result[n-1]
				result[n-1] = '_'
				result = append(result, prev, ch)
			case ChIsNextOfSepMark, ChIsNextOfKeepedMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOthers
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			if flag == ChIsNextOfSepMark {
				result = append(result, '_')
			}
			flag = ChIsNextOfKeepedMark
			result = append(result, ch)
		} else {
			if flag != ChIsFirstOfStr {
				flag = ChIsNextOfSepMark
			}
		}
	}

	return string(result)
}
