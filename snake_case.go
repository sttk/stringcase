// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

func SnakeCase(input string) string {
	return SnakeCaseWithNumsAsTail(input)
}

func SnakeCaseWithNumsAsHead(input string) string {
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
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
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

func SnakeCaseWithNumsAsTail(input string) string {
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

func SnakeCaseWithNumsAsWord(input string) string {
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
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
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

func SnakeCaseWithNumsAsPart(input string) string {
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
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
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

func SnakeCaseWithSep(input, seps string) string {
	return SnakeCaseWithNonAlphasAsTailAndSep(input, seps)
}

func SnakeCaseWithNonAlphasAsHeadAndSep(input, seps string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else {
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
			}
			flag = ChIsNextOfKeepedMark
		}
	}

	return string(result)
}

func SnakeCaseWithNonAlphasAsTailAndSep(input, seps string) string {
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

func SnakeCaseWithNonAlphasAsWordAndSep(input, seps string) string {
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
		} else {
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
			}
			flag = ChIsNextOfKeepedMark
		}
	}

	return string(result)
}

func SnakeCaseWithNonAlphasAsPartAndSep(input, seps string) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
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
	return SnakeCaseWithNonAlphasAsTailAndKeep(input, keeped)
}

func SnakeCaseWithNonAlphasAsHeadAndKeep(input, keeped string) string {
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
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
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

func SnakeCaseWithNonAlphasAsTailAndKeep(input, keeped string) string {
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

func SnakeCaseWithNonAlphasAsWordAndKeep(input, keeped string) string {
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
		} else if isAsciiDigit(ch) || strings.ContainsRune(keeped, ch) {
			switch flag {
			case ChIsFirstOfStr, ChIsNextOfKeepedMark:
				result = append(result, ch)
			default:
				result = append(result, '_', ch)
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

func SnakeCaseWithNonAlphasAsPartAndKeep(input, keeped string) string {
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
			case ChIsNextOfUpper, ChIsNextOfContdUpper, ChIsNextOfKeepedMark:
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
			case ChIsNextOfSepMark:
				result = append(result, '_', ch)
			default:
				result = append(result, ch)
			}
			flag = ChIsOther
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
