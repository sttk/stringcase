// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

// Options is a struct that represents options for case conversion of strings.
//
// The SeparateBeforeNonAlphabets field specifies whether to treat the
// beginning of a sequence of non-alphabetical characters as a word
// boundary. The SeparateAfterNonAlphabets field specifies whether to
// treat the end of a sequence of non-alphabetical characters as a word
// boundary. The Separators field specifies the set of characters to be
// treated as word separators and removed from the result string. The
// Keep field specifies the set of characters not to be treated as word
// separators and kept in the result string.
//
// Alphanumeric characters specified in Separators and Keep are ignored.
// If both Separators and Keep are specified, Separators takes precedence
// and Keep is ignored.
type Options struct {
	SeparateBeforeNonAlphabets bool
	SeparateAfterNonAlphabets  bool
	Separators                 string
	Keep                       string
}
