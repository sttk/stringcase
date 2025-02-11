package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleCamelCase() {
	camel := stringcase.CamelCase("foo_bar_baz")
	fmt.Printf("(1) camel = %s\n", camel)

	camel = stringcase.CamelCase("foo-Bar100baz")
	fmt.Printf("(2) camel = %s\n", camel)
	// Output:
	// (1) camel = fooBarBaz
	// (2) camel = fooBar100Baz
}

func ExampleCamelCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	camel := stringcase.CamelCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) camel = %s\n\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) camel = %s\n\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) camel = %s\n", camel)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	camel = stringcase.CamelCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) camel = %s\n", camel)
	// Output:
	// (1) camel = fooBar100Baz
	// (2) camel = fooBar100Baz
	// (3) camel = fooBar100baz
	// (4) camel = fooBar100baz
	//
	// (5) camel = fooBar100%Baz
	// (6) camel = fooBar100%Baz
	// (7) camel = fooBar100%baz
	// (8) camel = fooBar100%baz
	//
	// (9) camel = fooBar100%Baz
	// (a) camel = fooBar100%Baz
	// (b) camel = fooBar100%baz
	// (c) camel = fooBar100%baz
}

func ExampleCamelCaseWithSep() {
	camel := stringcase.CamelCaseWithSep("foo-bar100%baz", "- ")
	fmt.Printf("camel = %s\n", camel)
	// Output:
	// camel = fooBar100%Baz
}

func ExampleCamelCaseWithKeep() {
	camel := stringcase.CamelCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("camel = %s\n", camel)
	// Output:
	// camel = fooBar100%Baz
}
