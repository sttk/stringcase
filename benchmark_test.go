package stringcase_test

import (
	"testing"

	"github.com/sttk/stringcase"
)

// camel case

func BenchmarkCamelCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CamelCase("foo-bar100%baz")
	}
}

func BenchmarkCamelCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkCamelCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithSep("foo-bar100%baz", "%")
	}
}

// cobol case

func BenchmarkCobolCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CobolCase("foo-bar100%baz")
	}
}

func BenchmarkCobolCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkCobolCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithKeep("foo-bar100%baz", "%")
	}
}

// kebab case

func BenchmarkKebabCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.KebabCase("foo-bar100%baz")
	}
}

func BenchmarkKebabCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkKebabCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithKeep("foo-bar100%baz", "%")
	}
}

// macro case

func BenchmarkMacroCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.MacroCase("foo-bar100%baz")
	}
}

func BenchmarkMacroCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkMacroCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithKeep("foo-bar100%baz", "%")
	}
}

// pascal case

func BenchmarkPascalCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.PascalCase("foo-bar100%baz")
	}
}

func BenchmarkPascalCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkPascalCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithKeep("foo-bar100%baz", "%")
	}
}

// snake case

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCase("foo-bar100%baz")
	}
}

func BenchmarkSnakeCase_withSep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}

func BenchmarkSnakeCase_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}

// train case

func BenchmarkTrainCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.TrainCase("foo-bar100%baz")
	}
}

func BenchmarkTrainCase_withSep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}

func BenchmarkTrainCase_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}


// snake case with options

func BenchmarkSnakeCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkSnakeCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithOptions("foo-bar100%baz", opts)
	}
}

// train case with options

func BenchmarkTrainCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkTrainCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithOptions("foo-bar100%baz", opts)
	}
}
