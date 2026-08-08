// Copyright (C) 2024-2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// CobolCaseWithOptions converts the input string to cobol case with the
// specified options.
func CobolCaseWithOptions(input string, opts Options) string {
	return Upperize(input, '-', opts)
}

// CobolCase converts the input string to cobol case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func CobolCase(input string) string {
	return Upperize(input, '-', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
