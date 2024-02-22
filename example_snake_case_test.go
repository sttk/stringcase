package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleSnakeCase() {
	snake := stringcase.SnakeCase("fooBarBaz")
	fmt.Printf("snake = %s\n", snake)
	// Output:
	// snake = foo_bar_baz
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
