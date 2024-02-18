package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExamplePascalCase() {
	pascal := stringcase.PascalCase("foo_bar_baz")
	fmt.Printf("pascal = %s\n", pascal)
	// Output:
	// pascal = FooBarBaz
}

func ExamplePascalCaseWithSep() {
	pascal := stringcase.PascalCaseWithSep("foo-Bar100%Baz", "- ")
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
