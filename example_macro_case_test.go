package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleMacroCase() {
	macro := stringcase.MacroCase("fooBarBaz")
	fmt.Printf("(1) macro = %s\n", macro)

	macro = stringcase.MacroCase("foo-Bar100baz")
	fmt.Printf("(2) macro = %s\n", macro)
	// Output:
	// (1) macro = FOO_BAR_BAZ
	// (2) macro = FOO_BAR100_BAZ
}

func ExampleMacroCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	macro := stringcase.MacroCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) macro = %s\n\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) macro = %s\n\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) macro = %s\n", macro)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	macro = stringcase.MacroCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) macro = %s\n", macro)
	// Output:
	// (1) macro = FOO_BAR100_BAZ
	// (2) macro = FOO_BAR_100_BAZ
	// (3) macro = FOO_BAR_100BAZ
	// (4) macro = FOO_BAR100BAZ
	//
	// (5) macro = FOO_BAR100%_BAZ
	// (6) macro = FOO_BAR_100%_BAZ
	// (7) macro = FOO_BAR_100%BAZ
	// (8) macro = FOO_BAR100%BAZ
	//
	// (9) macro = FOO_BAR100%_BAZ
	// (a) macro = FOO_BAR_100%_BAZ
	// (b) macro = FOO_BAR_100%BAZ
	// (c) macro = FOO_BAR100%BAZ
}

func ExampleMacroCaseWithSep() {
	macro := stringcase.MacroCaseWithSep("foo-bar100%baz", "- ")
	fmt.Printf("macro = %s\n", macro)
	// Output:
	// macro = FOO_BAR100%_BAZ
}

func ExampleMacroCaseWithKeep() {
	macro := stringcase.MacroCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("macro = %s\n", macro)
	// Output:
	// macro = FOO_BAR100%_BAZ
}
