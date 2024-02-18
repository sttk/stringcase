package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleCobolCase() {
	cobol := stringcase.CobolCase("foo_bar_baz")
	fmt.Printf("cobol = %s\n", cobol)
	// Output:
	// cobol = FOO-BAR-BAZ
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
