// Copyright (C) 2024-2025 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// TrainCaseWithOptions converts the input string to train case with the
// specified options.
func TrainCaseWithOptions(input string, opts Options) string {
	result := make([]rune, 0, len(input)+len(input)/2)

	const (
		ChIsFirstOfStr = iota
		ChIsNextOfUpper
		ChIsNextOfContdUpper
		ChIsNextOfSepMark
		ChIsNextOfKeptMark
		ChIsOther
	)
	var flag uint8 = ChIsFirstOfStr

	for _, ch := range input {
		if isAsciiUpperCase(ch) {
			if flag == ChIsFirstOfStr {
				result = append(result, ch)
				flag = ChIsNextOfUpper
			} else if flag == ChIsNextOfUpper || flag == ChIsNextOfContdUpper ||
				(!opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			} else {
				result = append(result, '-', ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			if flag == ChIsFirstOfStr {
				result = append(result, toAsciiUpperCase(ch))
			} else if flag == ChIsNextOfContdUpper {
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					prev = toAsciiUpperCase(prev)
				}
				result[n-1] = '-'
				result = append(result, prev, ch)
			} else if flag == ChIsNextOfSepMark ||
				(opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, '-', toAsciiUpperCase(ch))
			} else {
				result = append(result, ch)
			}
			flag = ChIsOther
		} else {
			isKeptChar := false
			if len(opts.Separators) > 0 {
				if !strings.ContainsRune(opts.Separators, ch) {
					isKeptChar = true
				}
			} else if len(opts.Keep) > 0 {
				if isAsciiDigit(ch) || strings.ContainsRune(opts.Keep, ch) {
					isKeptChar = true
				}
			} else {
				if isAsciiDigit(ch) {
					isKeptChar = true
				}
			}

			if isKeptChar {
				if opts.SeparateBeforeNonAlphabets {
					if flag == ChIsFirstOfStr || flag == ChIsNextOfKeptMark {
						result = append(result, ch)
					} else {
						result = append(result, '-', ch)
					}
				} else {
					if flag != ChIsNextOfSepMark {
						result = append(result, ch)
					} else {
						result = append(result, '-', ch)
					}
				}
				flag = ChIsNextOfKeptMark
			} else {
				if flag != ChIsFirstOfStr {
					flag = ChIsNextOfSepMark
				}
			}
		}
	}

	return string(result)
}

// TrainCase converts the input string to train case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func TrainCase(input string) string {
	return TrainCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}

// TrainCaseWithSep converts the input string to train case with the
// specified separator characters.
func TrainCaseWithSep(input string, seps string) string {
	return TrainCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 seps,
	})
}

// TrainCaseWithKeep converts the input string to train case with the
// specified characters to be kept.
func TrainCaseWithKeep(input string, kept string) string {
	return TrainCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       kept,
	})
}
