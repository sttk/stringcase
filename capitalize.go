// Copyright (C) 2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

import (
	"strings"
)

// Capitalize converts the input string by capitalizing the first ASCII letter of each word and
// lowercasing subsequent letters, inserting the specified joiner rune between word boundaries
// according to the given options. It serves as a core engine for transforming input strings into
// capitalized casing styles, such as Train-Case or PascalCase, using custom joiner runes and
// customizable word separation rules defined in Options.
//
// During conversion, the initial ASCII letter of each word is converted to ASCII uppercase, while
// subsequent ASCII letters in that word are converted to ASCII lowercase. Word boundaries are
// automatically recognized at casing transitions, such as between lowercase and uppercase
// letters or before the final uppercase letter of an acronym preceding a lowercase sequence.
// When non-alphanumeric characters are encountered, ASCII digits are kept by default, while other
// characters are evaluated against Options. If opts.Separators is non-empty, characters matching
// opts.Separators are removed as separators while other non-alphanumeric characters are kept;
// otherwise, if opts.Keep is non-empty, specified characters are kept and all other
// non-alphanumeric characters are removed. If neither is specified, all non-alphanumeric
// characters are treated as separators and removed. The fields opts.SeparateBeforeNonAlphabets
// and opts.SeparateAfterNonAlphabets further determine whether word boundaries are inserted before
// or after non-alphabetic sequences.
//
// This function never returns an error or panics on any input, returning an empty string when the
// input is empty. Casing transformations and word boundary detections apply strictly to ASCII
// letters, treating non-ASCII characters as non-alphanumeric. If both opts.Separators and
// opts.Keep are specified, opts.Separators takes precedence and opts.Keep is ignored, while any
// alphanumeric characters listed in either field are disregarded. Additionally, leading and
// trailing separator characters are trimmed from the result without producing leading or trailing
// joiners.
func Capitalize(input string, joiner rune, opts Options) string {
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
				result = append(result, joiner, ch)
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
				result[n-1] = joiner
				result = append(result, prev, ch)
			} else if flag == ChIsNextOfSepMark ||
				(opts.SeparateAfterNonAlphabets && flag == ChIsNextOfKeptMark) {
				result = append(result, joiner, toAsciiUpperCase(ch))
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
						result = append(result, joiner, ch)
					}
				} else {
					if flag != ChIsNextOfSepMark {
						result = append(result, ch)
					} else {
						result = append(result, joiner, ch)
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
