package stringcase_test

import (
	"fmt"

	"github.com/sttk/stringcase"
)

func ExampleTrainCase() {
	train := stringcase.TrainCase("fooBarBaz")
	fmt.Printf("train = %s\n", train)
	// Output:
	// train = Foo-Bar-Baz
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
