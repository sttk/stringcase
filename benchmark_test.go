package stringcase_test

import (
	"testing"

	"github.com/sttk/stringcase"
)

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

func BenchmarkSnakeCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCase("foo-bar100%baz")
	}
}

func BenchmarkSnakeCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkSnakeCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.SnakeCaseWithKeep("foo-bar100%baz", "%")
	}
}

func BenchmarkTrainCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.TrainCase("foo-bar100%baz")
	}
}

func BenchmarkTrainCaseWithSep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithSep("foo-bar100%baz", "-")
	}
}

func BenchmarkTrainCaseWithKeep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringcase.TrainCaseWithKeep("foo-bar100%baz", "%")
	}
}
