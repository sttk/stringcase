package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestMacroCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.MacroCase("abcDefGHIjk")
		assert.Equal(t, result, "ABC_DEF_GH_IJK")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.MacroCase("AbcDefGHIjk")
		assert.Equal(t, result, "ABC_DEF_GH_IJK")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.MacroCase("abc_def_ghi")
		assert.Equal(t, result, "ABC_DEF_GHI")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.MacroCase("abc-def-ghi")
		assert.Equal(t, result, "ABC_DEF_GHI")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.MacroCase("Abc-Def-Ghi")
		assert.Equal(t, result, "ABC_DEF_GHI")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.MacroCase("ABC_DEF_GHI")
		assert.Equal(t, result, "ABC_DEF_GHI")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.MacroCase("ABC-DEF-GHI")
		assert.Equal(t, result, "ABC_DEF_GHI")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.MacroCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "ABC123_456_DEF_G89_HI_JKL_MN12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.MacroCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.MacroCase("123abc456def")
		assert.Equal(t, result, "123_ABC456_DEF")

		result = stringcase.MacroCase("123ABC456DEF")
		assert.Equal(t, result, "123_ABC456_DEF")

		result = stringcase.MacroCase("123Abc456Def")
		assert.Equal(t, result, "123_ABC456_DEF")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.MacroCase("")
		assert.Equal(t, result, "")
	})
}

func TestMacroCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456DEF_G_89HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456_DEF_G89_HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456_DEF_G_89_HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456DEF_G89HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456DEF_G_89HI_JKL_MN_12")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123-456DEF_G_89HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC_~!_DEF_#_GHI_%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456_DEF_G89_HI_JKL_MN12")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456_DEF_G89_HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._ABC~!_DEF#_GHI%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC___DEF___GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC___DEF___GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456_DEF_G_89_HI_JKL_MN_12")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123-456_DEF_G_89_HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._ABC_~!_DEF_#_GHI_%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "-"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456DEF_G89HI_JKL_MN12")

			opts.Separators = "_"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF_G89HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC~!_DEF#_GHI%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_-DEF_-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456DEF_G_89HI_JKL_MN_12")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123-456DEF_G_89HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC_~!_DEF_#_GHI_%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC_456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC__DEF__GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456_DEF_G89_HI_JKL_MN12")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456_DEF_G89_HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._ABC~!_DEF#_GHI%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC___DEF___GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC___DEF___GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_-_DEF_-_GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123_456_DEF_G_89_HI_JKL_MN_12")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC_123-456_DEF_G_89_HI_JKL_MN_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._ABC_~!_DEF_#_GHI_%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC_456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
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
			result := stringcase.MacroCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "ABC_DEF_GH_IJK")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "ABC-_DEF-_GHI")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "_"
			result = stringcase.MacroCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC_DEF_GHI")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "ABC-DEF-GHI")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123_456DEF_G89HI_JKL_MN12")

			opts.Keep = "-"
			result = stringcase.MacroCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "ABC123-456DEF_G89HI_JKL_MN12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.MacroCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".ABC~!_DEF#_GHI%_JK_LM_NO_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123ABC456DEF")

			result = stringcase.MacroCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_ABC456_DEF")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.MacroCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
