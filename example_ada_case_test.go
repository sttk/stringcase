package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleAdaCase() {
	ada := stringcase.AdaCase("fooBarBaz")
	fmt.Printf("(1) ada = %s\n", ada)

	ada = stringcase.AdaCase("foo-Bar100baz")
	fmt.Printf("(2) ada = %s\n", ada)
	// Output:
	// (1) ada = Foo_Bar_Baz
	// (2) ada = Foo_Bar100_Baz
}

func ExampleAdaCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	ada := stringcase.AdaCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) ada = %s\n\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) ada = %s\n\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) ada = %s\n", ada)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	ada = stringcase.AdaCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) ada = %s\n", ada)
	// Output:
	// (1) ada = Foo_Bar100_Baz
	// (2) ada = Foo_Bar_100_Baz
	// (3) ada = Foo_Bar_100baz
	// (4) ada = Foo_Bar100baz
	//
	// (5) ada = Foo_Bar100%_Baz
	// (6) ada = Foo_Bar_100%_Baz
	// (7) ada = Foo_Bar_100%baz
	// (8) ada = Foo_Bar100%baz
	//
	// (9) ada = Foo_Bar100%_Baz
	// (a) ada = Foo_Bar_100%_Baz
	// (b) ada = Foo_Bar_100%baz
	// (c) ada = Foo_Bar100%baz
}
