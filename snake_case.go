// Copyright (C) 2024-2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// SnakeCaseWithOptions converts the input string to snake case with the
// specified options.
func SnakeCaseWithOptions(input string, opts Options) string {
	return Lowerize(input, '_', opts)
}

// SnakeCase converts the input string to snake case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func SnakeCase(input string) string {
	return Lowerize(input, '_', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
