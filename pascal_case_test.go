package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestPascalCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.PascalCase("abcDefGHIjk")
		assert.Equal(t, result, "AbcDefGhIjk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.PascalCase("AbcDefGHIjk")
		assert.Equal(t, result, "AbcDefGhIjk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.PascalCase("abc_def_ghi")
		assert.Equal(t, result, "AbcDefGhi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.PascalCase("abc-def-ghi")
		assert.Equal(t, result, "AbcDefGhi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.PascalCase("Abc-Def-Ghi")
		assert.Equal(t, result, "AbcDefGhi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.PascalCase("ABC_DEF_GHI")
		assert.Equal(t, result, "AbcDefGhi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.PascalCase("ABC-DEF-GHI")
		assert.Equal(t, result, "AbcDefGhi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.PascalCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "Abc123456DefG89HiJklMn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.PascalCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "AbcDefGhiJkLmNo")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.PascalCase("123abc456def")
		assert.Equal(t, result, "123Abc456Def")

		result = stringcase.PascalCase("123ABC456DEF")
		assert.Equal(t, result, "123Abc456Def")

		result = stringcase.PascalCase("123Abc456Def")
		assert.Equal(t, result, "123Abc456Def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.PascalCase("")
		assert.Equal(t, result, "")
	})
}

func TestPascalCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "AbcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "AbcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "AbcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "AbcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as head of a word and with separators", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.PascalCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123def")
		})
	})

	t.Run("non-alphabets as tail of a word and with separators", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.PascalCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123Def")
		})
	})

	t.Run("non-alphabets as a word and with separators", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.PascalCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123Def")
		})
	})

	t.Run("non-alphabets as part of a word and with separators", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "-"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")

			opts.Separators = "_"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.PascalCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123def")
		})
	})

	t.Run("non-alphabets as head of a word and with kept characters", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word and with kept characters", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word and with kept characters", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456DefG89HiJklMn12")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word and with kept characters", func(t *testing.T) {
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "AbcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "_"
			result = stringcase.PascalCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "AbcDefGhi")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123456defG89hiJklMn12")

			opts.Keep = "-"
			result = stringcase.PascalCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.PascalCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.PascalCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.PascalCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
