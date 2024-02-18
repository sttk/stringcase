// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
This library provides some functions that convert string cases between
camelCase, COBOL-CASE, kebab-case, MACRO_CASE, PascalCase, snake_case and
Train-Case.

Basically, these functions targets the upper and lower cases of only ASCII
alphabets for capitalization, and all characters except ASCII alphabets and
ASCII numbers are eliminated as word separators.

To limit characters using as separators, the functions named like *_with_sep
are provided, and to keep specified characters, the functions named like
*_with_keep are provided.
*/
package stringcase
