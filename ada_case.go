// Copyright (C) 2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// AdaCaseWithOptions converts the input string to Ada case with the
// specified options.
func AdaCaseWithOptions(input string, opts Options) string {
	return Capitalize(input, '_', opts)
}

// AdaCase converts the input string to ada case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func AdaCase(input string) string {
	return Capitalize(input, '_', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
