// Copyright (C) 2024-2025 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// TitleCaseWithOptions converts the input string to title case with the
// specified options.
func TitleCaseWithOptions(input string, opts Options) string {
	return Capitalize(input, ' ', opts)
}

// TitleCase converts the input string to title case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func TitleCase(input string) string {
	return Capitalize(input, ' ', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
