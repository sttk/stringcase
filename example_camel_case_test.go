package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleCamelCase() {
	camel := stringcase.CamelCase("foo_bar_baz")
	fmt.Printf("camel = %s\n", camel)
	// Output:
	// camel = fooBarBaz
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
