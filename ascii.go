// Copyright (C) 2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package stringcase

func isAsciiUpperCase(r rune) bool {
	return (0x41 <= r && r <= 0x5a)
}

func isAsciiLowerCase(r rune) bool {
	return (0x61 <= r && r <= 0x7a)
}

func isAsciiDigit(r rune) bool {
	return (0x30 <= r && r <= 0x39)
}

func toAsciiUpperCase(r rune) rune {
	return (r + 0x41 - 0x61)
}

func toAsciiLowerCase(r rune) rune {
	return (r + 0x61 - 0x41)
}
