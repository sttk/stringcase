package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleTitleCase() {
	title := stringcase.TitleCase("fooBarBaz")
	fmt.Printf("(1) title = %s\n", title)

	title = stringcase.TitleCase("foo-Bar100baz")
	fmt.Printf("(2) title = %s\n", title)
	// Output:
	// (1) title = Foo Bar Baz
	// (2) title = Foo Bar100 Baz
}

func ExampleTitleCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	title := stringcase.TitleCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	title = stringcase.TitleCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	title = stringcase.TitleCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	title = stringcase.TitleCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) title = %s\n\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) title = %s\n\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) title = %s\n", title)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	title = stringcase.TitleCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) title = %s\n", title)
	// Output:
	// (1) title = Foo Bar100 Baz
	// (2) title = Foo Bar 100 Baz
	// (3) title = Foo Bar 100baz
	// (4) title = Foo Bar100baz
	//
	// (5) title = Foo Bar100% Baz
	// (6) title = Foo Bar 100% Baz
	// (7) title = Foo Bar 100%baz
	// (8) title = Foo Bar100%baz
	//
	// (9) title = Foo Bar100% Baz
	// (a) title = Foo Bar 100% Baz
	// (b) title = Foo Bar 100%baz
	// (c) title = Foo Bar100%baz
}
