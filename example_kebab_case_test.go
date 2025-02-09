package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleKebabCase() {
	kebab := stringcase.KebabCase("fooBarBaz")
	fmt.Printf("(1) kebab = %s\n", kebab)

	kebab = stringcase.KebabCase("foo-Bar100baz")
	fmt.Printf("(2) kebab = %s\n", kebab)
	// Output:
	// (1) kebab = foo-bar-baz
	// (2) kebab = foo-bar100-baz
}

func ExampleKebabCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	kebab := stringcase.KebabCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) kebab = %s\n\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) kebab = %s\n\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) kebab = %s\n", kebab)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	kebab = stringcase.KebabCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) kebab = %s\n", kebab)
	// Output:
	// (1) kebab = foo-bar100-baz
	// (2) kebab = foo-bar-100-baz
	// (3) kebab = foo-bar-100baz
	// (4) kebab = foo-bar100baz
	//
	// (5) kebab = foo-bar100%-baz
	// (6) kebab = foo-bar-100%-baz
	// (7) kebab = foo-bar-100%baz
	// (8) kebab = foo-bar100%baz
	//
	// (9) kebab = foo-bar100%-baz
	// (a) kebab = foo-bar-100%-baz
	// (b) kebab = foo-bar-100%baz
	// (c) kebab = foo-bar100%baz
}

func ExampleKebabCaseWithSep() {
	kebab := stringcase.KebabCaseWithSep("foo-Bar100%Baz", "- ")
	fmt.Printf("kebab = %s\n", kebab)
	// Output:
	// kebab = foo-bar100%-baz
}

func ExampleKebabCaseWithKeep() {
	kebab := stringcase.KebabCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("kebab = %s\n", kebab)
	// Output:
	// kebab = foo-bar100%-baz
}
