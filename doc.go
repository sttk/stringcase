// Copyright (C) 2024-2026 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
This library provides some functions that convert string cases between Ada_Case, camelCase,
COBOL-CASE, kebab-case, MACRO_CASE, PascalCase, snake_case, Title Case, and Train-Case.
In addition, the functions Capitalize, Lowerize, and Upperize are provided to convert
string cases with a custom joiner character.

Essentially, these functions only target ASCII uppercase and lowercase letters for capitalization.
All characters other than ASCII uppercase and lowercase letters and ASCII numbers are removed as
word separators.

If you want to use some symbols as separators, specify those symbols in the Separators field of
Options struct and use the 〜CaseWithOptions function for the desired case.
If you want to retain certain symbols and use everything else as separators, specify those symbols
in Keep field of Options struct and use the 〜CaseWithOptions function for the desired case.

Additionally, you can specify whether to place word boundaries before and/or after non-alphabetic
characters with conversion options.
This can be set using the SeparateBeforeNonAlphabets and SeparateAfterNonAlphabets fields in the
Options struct.

The 〜Case functions that do not take Options as an argument only place word boundaries after
non-alphabetic characters.
In other words, they behave as if
SeparateBeforeNonAlphabets = false and SeparateAfterNonAlphabets = true.
*/
package stringcase
