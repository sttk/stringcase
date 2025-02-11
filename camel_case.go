// Copyright (C) 2024-2025 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// CamelCaseWithOptions converts the input string to pascal case with the
// specified options.
func CamelCaseWithOptions(input string, opts Options) string {
	result := make([]rune, 0, len(input))

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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			} else if flag == ChIsNextOfUpper || flag == ChIsNextOfContdUpper ||
				(!opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			} else {
				result = append(result, ch)
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			if flag == ChIsNextOfContdUpper {
				n := len(result)
				prev := result[n-1]
				if isAsciiLowerCase(prev) {
					prev = toAsciiUpperCase(prev)
				}
				result[n-1] = prev
				result = append(result, ch)
			} else if flag == ChIsNextOfSepMark ||
				(opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, toAsciiUpperCase(ch))
			} else {
				result = append(result, ch)
			}
			flag = ChIsOther
		} else {
			isKeptChar := false
			if isAsciiDigit(ch) {
				isKeptChar = true
			} else if len(opts.Separators) > 0 {
				if !strings.ContainsRune(opts.Separators, ch) {
					isKeptChar = true
				}
			} else if len(opts.Keep) > 0 {
				if strings.ContainsRune(opts.Keep, ch) {
					isKeptChar = true
				}
			}

			if isKeptChar {
				result = append(result, ch)
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

// CamelCase converts the input string to pascal case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func CamelCase(input string) string {
	return CamelCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}

// CamelCaseWithSep converts the input string to pascal case with the
// specified separator characters.
//
// Deprecated: Should use CamelCaseWithOptions instead
func CamelCaseWithSep(input string, seps string) string {
	return CamelCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 seps,
	})
}

// CamelCaseWithKeep converts the input string to pascal case with the
// specified characters to be kept.
//
// Deprecated: Should use CamelCaseWithOptions instead
func CamelCaseWithKeep(input string, kept string) string {
	return CamelCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       kept,
	})
}
