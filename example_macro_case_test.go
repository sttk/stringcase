package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleMacroCase() {
	macro := stringcase.MacroCase("fooBarBaz")
	fmt.Printf("macro = %s\n", macro)
	// Output:
	// macro = FOO_BAR_BAZ
}

func ExampleMacroCaseWithSep() {
	macro := stringcase.MacroCaseWithSep("foo-Bar100%Baz", "- ")
	fmt.Printf("macro = %s\n", macro)
	// Output:
	// macro = FOO_BAR100%_BAZ
}

func ExampleMacroCaseWithKeep() {
	macro := stringcase.MacroCaseWithKeep("foo-Bar100%Baz", "%")
	fmt.Printf("macro = %s\n", macro)
	// Output:
	// macro = FOO_BAR100%_BAZ
}
