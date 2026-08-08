// Copyright (C) 2024-2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// KebabCaseWithOptions converts the input string to kebab case with the
// specified options.
func KebabCaseWithOptions(input string, opts Options) string {
	return Lowerize(input, '-', opts)
}

// KebabCase converts the input string to kebab case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func KebabCase(input string) string {
	return Lowerize(input, '-', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
