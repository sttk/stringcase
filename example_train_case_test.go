package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleTrainCase() {
	train := stringcase.TrainCase("fooBarBaz")
	fmt.Printf("(1) train = %s\n", train)

	train = stringcase.TrainCase("foo-Bar100baz")
	fmt.Printf("(2) train = %s\n", train)
	// Output:
	// (1) train = Foo-Bar-Baz
	// (2) train = Foo-Bar100-Baz
}

func ExampleTrainCaseWithOptions() {
	opts := stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true}
	train := stringcase.TrainCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(1) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true}
	train = stringcase.TrainCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(2) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false}
	train = stringcase.TrainCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(3) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false}
	train = stringcase.TrainCaseWithOptions("foo#Bar100baz", opts)
	fmt.Printf("(4) train = %s\n\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Separators: "#"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(5) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Separators: "#"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(6) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Separators: "#"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(7) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Separators: "#"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(8) train = %s\n\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: true, Keep: "%"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(9) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: true, Keep: "%"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(a) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: true, SeparateAfterNonAlphabets: false, Keep: "%"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(b) train = %s\n", train)

	opts = stringcase.Options{SeparateBeforeNonAlphabets: false, SeparateAfterNonAlphabets: false, Keep: "%"}
	train = stringcase.TrainCaseWithOptions("foo#Bar100%baz", opts)
	fmt.Printf("(c) train = %s\n", train)
	// Output:
	// (1) train = Foo-Bar100-Baz
	// (2) train = Foo-Bar-100-Baz
	// (3) train = Foo-Bar-100baz
	// (4) train = Foo-Bar100baz
	//
	// (5) train = Foo-Bar100%-Baz
	// (6) train = Foo-Bar-100%-Baz
	// (7) train = Foo-Bar-100%baz
	// (8) train = Foo-Bar100%baz
	//
	// (9) train = Foo-Bar100%-Baz
	// (a) train = Foo-Bar-100%-Baz
	// (b) train = Foo-Bar-100%baz
	// (c) train = Foo-Bar100%baz
}

func ExampleTrainCaseWithSep() {
	train := stringcase.TrainCaseWithSep("foo-bar100%baz", "- ")
	fmt.Printf("train = %s\n", train)
	// Output:
	// train = Foo-Bar100%-Baz
}

func ExampleTrainCaseWithKeep() {
	train := stringcase.TrainCaseWithKeep("foo-bar100%baz", "%")
	fmt.Printf("train = %s\n", train)
	// Output:
	// train = Foo-Bar100%-Baz
}
