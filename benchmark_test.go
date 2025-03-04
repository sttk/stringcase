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

// camel case with options

func BenchmarkCamelCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCamelCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CamelCaseWithOptions("foo-bar100%baz", opts)
	}
}

// cobol case with options

func BenchmarkCobolCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkCobolCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.CobolCaseWithOptions("foo-bar100%baz", opts)
	}
}

// kebab case with options

func BenchmarkKebabCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkKebabCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.KebabCaseWithOptions("foo-bar100%baz", opts)
	}
}

// macro case with options

func BenchmarkMacroCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkMacroCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.MacroCaseWithOptions("foo-bar100%baz", opts)
	}
}

// pascal case with options

func BenchmarkPascalCase_nonAlphabetsAsHead(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsTail(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsWord(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsPart(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsHead_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsTail_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsWord_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsPart_withSeparators(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Separators:                 "-",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsHead_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsTail_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsWord_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: true,
		SeparateAfterNonAlphabets:  true,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
	}
}
func BenchmarkPascalCase_nonAlphabetsAsPart_withKeep(b *testing.B) {
	opts := stringcase.Options{
		SeparateBeforeNonAlphabets: false,
		SeparateAfterNonAlphabets:  false,
		Keep:                       "%",
	}
	for i := 0; i < b.N; i++ {
		stringcase.PascalCaseWithOptions("foo-bar100%baz", opts)
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
