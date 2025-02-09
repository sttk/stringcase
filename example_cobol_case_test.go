package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleCobolCase() {
	cobol := stringcase.CobolCase("fooBarBaz")
	fmt.Printf("(1) cobol = %s\n", cobol)

	cobol = stringcase.CobolCase("foo-Bar100baz")
	fmt.Printf("(2) cobol = %s\n", cobol)
	// Output:
	// (1) cobol = FOO-BAR-BAZ
	// (2) cobol = FOO-BAR100-BAZ
}

func ExampleCobolCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	cobol := stringcase.CobolCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) cobol = %s\n\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) cobol = %s\n\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) cobol = %s\n", cobol)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	cobol = stringcase.CobolCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) cobol = %s\n", cobol)
	// Output:
	// (1) cobol = FOO-BAR100-BAZ
	// (2) cobol = FOO-BAR-100-BAZ
	// (3) cobol = FOO-BAR-100BAZ
	// (4) cobol = FOO-BAR100BAZ
	//
	// (5) cobol = FOO-BAR100%-BAZ
	// (6) cobol = FOO-BAR-100%-BAZ
	// (7) cobol = FOO-BAR-100%BAZ
	// (8) cobol = FOO-BAR100%BAZ
	//
	// (9) cobol = FOO-BAR100%-BAZ
	// (a) cobol = FOO-BAR-100%-BAZ
	// (b) cobol = FOO-BAR-100%BAZ
	// (c) cobol = FOO-BAR100%BAZ
}

func ExampleCobolCaseWithSep() {
	cobol := stringcase.CobolCaseWithSep("foo-bar100%baz", "- ")
	fmt.Printf("cobol = %s\n", cobol)
	// Output:
	// cobol = FOO-BAR100%-BAZ
}

func ExampleCobolCaseWithKeep() {
	cobol := stringcase.CobolCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("cobol = %s\n", cobol)
	// Output:
	// cobol = FOO-BAR100%-BAZ
}
