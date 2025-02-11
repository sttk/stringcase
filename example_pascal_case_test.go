package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExamplePascalCase() {
	pascal := stringcase.PascalCase("foo_bar_baz")
	fmt.Printf("(1) pascal = %s\n", pascal)

	pascal = stringcase.PascalCase("foo-Bar100baz")
	fmt.Printf("(2) pascal = %s\n", pascal)
	// Output:
	// (1) pascal = FooBarBaz
	// (2) pascal = FooBar100Baz
}

func ExamplePascalCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	pascal := stringcase.PascalCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) pascal = %s\n\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) pascal = %s\n\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) pascal = %s\n", pascal)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	pascal = stringcase.PascalCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) pascal = %s\n", pascal)
	// Output:
	// (1) pascal = FooBar100Baz
	// (2) pascal = FooBar100Baz
	// (3) pascal = FooBar100baz
	// (4) pascal = FooBar100baz
	//
	// (5) pascal = FooBar100%Baz
	// (6) pascal = FooBar100%Baz
	// (7) pascal = FooBar100%baz
	// (8) pascal = FooBar100%baz
	//
	// (9) pascal = FooBar100%Baz
	// (a) pascal = FooBar100%Baz
	// (b) pascal = FooBar100%baz
	// (c) pascal = FooBar100%baz
}

func ExamplePascalCaseWithSep() {
	pascal := stringcase.PascalCaseWithSep("foo-bar100%baz", "- ")
	fmt.Printf("pascal = %s\n", pascal)
	// Output:
	// pascal = FooBar100%Baz
}

func ExamplePascalCaseWithKeep() {
	pascal := stringcase.PascalCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("pascal = %s\n", pascal)
	// Output:
	// pascal = FooBar100%Baz
}
