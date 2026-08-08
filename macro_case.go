// Copyright (C) 2024-2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// MacroCaseWithOptions converts the input string to macro case with the
// specified options.
func MacroCaseWithOptions(input string, opts Options) string {
	return Upperize(input, '_', opts)
}

// MacroCase converts the input string to macro case.
//
// It treats the end of a sequence of non-alphabetical characters as a
// word boundary, but not the beginning.
func MacroCase(input string) string {
	return Upperize(input, '_', Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	})
}
