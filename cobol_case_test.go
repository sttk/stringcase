package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestCobolCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.CobolCase("abcDefGHIjk")
		assert.Equal(t, result, "ABC-DEF-GH-IJK")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.CobolCase("AbcDefGHIjk")
		assert.Equal(t, result, "ABC-DEF-GH-IJK")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.CobolCase("abc_def_ghi")
		assert.Equal(t, result, "ABC-DEF-GHI")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.CobolCase("abc-def-ghi")
		assert.Equal(t, result, "ABC-DEF-GHI")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.CobolCase("Abc-Def-Ghi")
		assert.Equal(t, result, "ABC-DEF-GHI")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.CobolCase("ABC_DEF_GHI")
		assert.Equal(t, result, "ABC-DEF-GHI")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.CobolCase("ABC-DEF-GHI")
		assert.Equal(t, result, "ABC-DEF-GHI")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.CobolCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.CobolCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.CobolCase("123abc456def")
		assert.Equal(t, result, "123-ABC456-DEF")

		result = stringcase.CobolCase("123ABC456DEF")
		assert.Equal(t, result, "123-ABC456-DEF")

		result = stringcase.CobolCase("123Abc456Def")
		assert.Equal(t, result, "123-ABC456-DEF")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.CobolCase("")
		assert.Equal(t, result, "")
	})
}

func TestCobolCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456DEF-G-89HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456-DEF-G-89-HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF-G89HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.CobolCaseWithOptions("", opts)
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456DEF-G-89HI-JKL-MN-12")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456DEF-G-89HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC-~!-DEF-#-GHI-%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-ABC~!-DEF#-GHI%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-_-DEF-_-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-_-DEF-_-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456-DEF-G-89-HI-JKL-MN-12")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456-DEF-G-89-HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-ABC-~!-DEF-#-GHI-%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "-"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF-G89HI-JKL-MN12")

			opts.Separators = "_"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF-G89HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC~!-DEF#-GHI%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456DEF-G-89HI-JKL-MN-12")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456DEF-G-89HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC-~!-DEF-#-GHI-%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC-456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456-DEF-G89-HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-ABC~!-DEF#-GHI%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-_-DEF-_-GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-_-DEF-_-GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC---DEF---GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456-DEF-G-89-HI-JKL-MN-12")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC-123-456-DEF-G-89-HI-JKL-MN-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-ABC-~!-DEF-#-GHI-%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC-456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
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
			result := stringcase.CobolCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC-DEF-GH-IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC--DEF--GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "_"
			result = stringcase.CobolCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF-G89HI-JKL-MN12")

			opts.Keep = "-"
			result = stringcase.CobolCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF-G89HI-JKL-MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.CobolCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC~!-DEF#-GHI%-JK-LM-NO-?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.CobolCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-ABC456-DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.CobolCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
