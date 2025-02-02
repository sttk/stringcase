package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleSnakeCase() {
	snake := stringcase.SnakeCase("fooBarBaz")
	fmt.Printf("(1) snake = %s\n", snake)

	snake = stringcase.SnakeCase("foo-Bar100baz")
	fmt.Printf("(2) snake = %s\n", snake)
	// Output:
	// (1) snake = foo_bar_baz
	// (2) snake = foo_bar100_baz
}

func ExampleSnakeCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	snake := stringcase.SnakeCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) snake = %s\n\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) snake = %s\n\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) snake = %s\n", snake)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	snake = stringcase.SnakeCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) snake = %s\n", snake)
	// Output:
	// (1) snake = foo_bar100_baz
	// (2) snake = foo_bar_100_baz
	// (3) snake = foo_bar_100baz
	// (4) snake = foo_bar100baz
	//
	// (5) snake = foo_bar100%_baz
	// (6) snake = foo_bar_100%_baz
	// (7) snake = foo_bar_100%baz
	// (8) snake = foo_bar100%baz
	//
	// (9) snake = foo_bar100%_baz
	// (a) snake = foo_bar_100%_baz
	// (b) snake = foo_bar_100%baz
	// (c) snake = foo_bar100%baz
}

func ExampleSnakeCaseWithSep() {
	snake := stringcase.SnakeCaseWithSep("foo-Bar100%Baz", "- ")
	fmt.Printf("snake = %s\n", snake)
	// Output:
	// snake = foo_bar100%_baz
}

func ExampleSnakeCaseWithKeep() {
	snake := stringcase.SnakeCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("snake = %s\n", snake)
	// Output:
	// snake = foo_bar100%_baz
}
