// Copyright (C) 2024-2025 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// KebabCaseWithOptions converts the input string to kebab case with the
// specified options.
func KebabCaseWithOptions(input string, opts Options) string {
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
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			} else if flag == ChIsNextOfUpper || flag == ChIsNextOfContdUpper ||
				(!opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, toAsciiLowerCase(ch))
				flag = ChIsNextOfContdUpper
			} else {
				result = append(result, '-', toAsciiLowerCase(ch))
				flag = ChIsNextOfUpper
			}
		} else if isAsciiLowerCase(ch) {
			if flag == ChIsNextOfContdUpper {
				n := len(result)
				prev := result[n-1]
				result[n-1] = '-'
				result = append(result, prev, ch)
			} else if flag == ChIsNextOfSepMark ||
				(opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, '-', ch)
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

// KebabCase converts the input string to kebab case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func KebabCase(input string) string {
	return KebabCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}

// KebabCaseWithSep converts the input string to kebab case with the
// specified separator characters.
//
// Deprecated: Should use KebabCaseWithOptions instead
func KebabCaseWithSep(input string, seps string) string {
	return KebabCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 seps,
	})
}

// KebabCaseWithSep converts the input string to kebab case with the
// specified separator characters.
//
// Deprecated: Should use KebabCaseWithOptions instead
func KebabCaseWithKeep(input string, kept string) string {
	return KebabCaseWithOptions(input, Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       kept,
	})
}
