package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleKebabCase() {
	kebab := stringcase.KebabCase("fooBarBaz")
	fmt.Printf("kebab = %s\n", kebab)
	// Output:
	// kebab = foo-bar-baz
}

func ExampleKebabCaseWithSep() {
	kebab := stringcase.KebabCaseWithSep("foo-bar100%baz", "- ")
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
