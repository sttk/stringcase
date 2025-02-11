package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestCamelCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.CamelCase("abcDefGHIjk")
		assert.Equal(t, result, "abcDefGhIjk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.CamelCase("AbcDefGHIjk")
		assert.Equal(t, result, "abcDefGhIjk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.CamelCase("abc_def_ghi")
		assert.Equal(t, result, "abcDefGhi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.CamelCase("abc-def-ghi")
		assert.Equal(t, result, "abcDefGhi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.CamelCase("Abc-Def-Ghi")
		assert.Equal(t, result, "abcDefGhi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.CamelCase("ABC_DEF_GHI")
		assert.Equal(t, result, "abcDefGhi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.CamelCase("ABC-DEF-GHI")
		assert.Equal(t, result, "abcDefGhi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.CamelCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc123456DefG89HiJklMn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.CamelCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abcDefGhiJkLmNo")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.CamelCase("123abc456def")
		assert.Equal(t, result, "123Abc456Def")

		result = stringcase.CamelCase("123ABC456DEF")
		assert.Equal(t, result, "123Abc456Def")

		result = stringcase.CamelCase("123Abc456Def")
		assert.Equal(t, result, "123Abc456Def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.CamelCase("")
		assert.Equal(t, result, "")
	})
}

func TestCamelCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abcDefGhiJkLmNo")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CamelCaseWithOptions("", opts)
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.CamelCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123def")
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.CamelCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123Def")
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.CamelCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123Def")
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "-"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")

			opts.Separators = "_"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.CamelCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123def")
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456DefG89HiJklMn12")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456DefG89HiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123Abc456Def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
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
			result := stringcase.CamelCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abcDefGhIjk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-Def-Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "_"
			result = stringcase.CamelCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abcDefGhi")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123456defG89hiJklMn12")

			opts.Keep = "-"
			result = stringcase.CamelCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456defG89hiJklMn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CamelCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!Def#Ghi%JkLmNo?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.CamelCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123Abc456Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CamelCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
