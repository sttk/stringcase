package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestAdaCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.AdaCase("abcDefGHIjk")
		assert.Equal(t, result, "Abc_Def_Gh_Ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.AdaCase("AbcDefGHIjk")
		assert.Equal(t, result, "Abc_Def_Gh_Ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.AdaCase("abc_def_ghi")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.AdaCase("abc-def-ghi")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.AdaCase("Abc-Def-Ghi")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert Title Case", func(t *testing.T) {
		result := stringcase.AdaCase("Abc Def Ghi")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert Ada_Case", func(t *testing.T) {
		result := stringcase.AdaCase("Abc_Def_Ghi")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.AdaCase("ABC_DEF_GHI")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.AdaCase("ABC-DEF-GHI")
		assert.Equal(t, result, "Abc_Def_Ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.AdaCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "Abc123_456_Def_G89_Hi_Jkl_Mn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.AdaCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "Abc_Def_Ghi_Jk_Lm_No")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.AdaCase("123abc456def")
		assert.Equal(t, result, "123_Abc456_Def")

		result = stringcase.AdaCase("123ABC456DEF")
		assert.Equal(t, result, "123_Abc456_Def")

		result = stringcase.AdaCase("123Abc456Def")
		assert.Equal(t, result, "123_Abc456_Def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.AdaCase("")
		assert.Equal(t, result, "")
	})
}

func TestAdaCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456def_G_89hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc_Def_Ghi_Jk_Lm_No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456_Def_G89_Hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc_Def_Ghi_Jk_Lm_No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456_Def_G_89_Hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc_Def_Ghi_Jk_Lm_No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456def_G89hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc_Def_Ghi_Jk_Lm_No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.AdaCaseWithOptions("", opts)
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_-def_-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_ _Def_ _Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_-def_-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456def_G_89hi_Jkl_Mn_12")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123-456def_G_89hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc_~!_Def_#_Ghi_%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.AdaCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc_123def")
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc _Def _Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456_Def_G89_Hi_Jkl_Mn12")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456_Def_G89_Hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._Abc~!_Def#_Ghi%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.AdaCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123_Def")
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_ _Def_ _Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456_Def_G_89_Hi_Jkl_Mn_12")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123-456_Def_G_89_Hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._Abc_~!_Def_#_Ghi_%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.AdaCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc_123_Def")
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc _Def _Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "-"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456def_G89hi_Jkl_Mn12")

			opts.Separators = "_"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456def_G89hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!_Def#_Ghi%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.AdaCaseWithOptions("abc123def", opts)
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_-def_-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = " "
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_ _Def_ _Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_-def_-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456def_G_89hi_Jkl_Mn_12")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123-456def_G_89hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc_~!_Def_#_Ghi_%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = " "
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc _Def _Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456_Def_G89_Hi_Jkl_Mn12")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456_Def_G89_Hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._Abc~!_Def#_Ghi%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = " "
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_ _Def_ _Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc___Def___Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_-_Def_-_Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123_456_Def_G_89_Hi_Jkl_Mn_12")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc_123-456_Def_G_89_Hi_Jkl_Mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._Abc_~!_Def_#_Ghi_%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_Abc_456_Def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc_456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
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
			result := stringcase.AdaCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc_Def_Gh_Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc-_Def-_Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = " "
			result = stringcase.AdaCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc _Def _Ghi")
		})

		t.Run("convert Ada-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc__Def__Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "_"
			result = stringcase.AdaCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc_Def_Ghi")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123_456def_G89hi_Jkl_Mn12")

			opts.Keep = "-"
			result = stringcase.AdaCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456def_G89hi_Jkl_Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.AdaCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!_Def#_Ghi%_Jk_Lm_No_?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.AdaCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_Abc456_Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.AdaCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
