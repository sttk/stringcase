package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestTitleCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.TitleCase("abcDefGHIjk")
		assert.Equal(t, result, "Abc Def Gh Ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.TitleCase("AbcDefGHIjk")
		assert.Equal(t, result, "Abc Def Gh Ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.TitleCase("abc_def_ghi")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.TitleCase("abc-def-ghi")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.TitleCase("Abc-Def-Ghi")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert Title Case", func(t *testing.T) {
		result := stringcase.TitleCase("Abc Def Ghi")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert Ada_Case", func(t *testing.T) {
		result := stringcase.TitleCase("Abc_Def_Ghi")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.TitleCase("ABC_DEF_GHI")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.TitleCase("ABC-DEF-GHI")
		assert.Equal(t, result, "Abc Def Ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.TitleCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "Abc123 456 Def G89 Hi Jkl Mn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.TitleCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "Abc Def Ghi Jk Lm No")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.TitleCase("123abc456def")
		assert.Equal(t, result, "123 Abc456 Def")

		result = stringcase.TitleCase("123ABC456DEF")
		assert.Equal(t, result, "123 Abc456 Def")

		result = stringcase.TitleCase("123Abc456Def")
		assert.Equal(t, result, "123 Abc456 Def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.TitleCase("")
		assert.Equal(t, result, "")
	})
}

func TestTitleCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456def G 89hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc Def Ghi Jk Lm No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456 Def G89 Hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc Def Ghi Jk Lm No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456 Def G 89 Hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc Def Ghi Jk Lm No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456def G89hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "Abc Def Ghi Jk Lm No")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.TitleCaseWithOptions("", opts)
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc _def _ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc -def -ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc   Def   Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc _def _ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc -def -ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456def G 89hi Jkl Mn 12")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123-456def G 89hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc ~! Def # Ghi % Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.TitleCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc 123def")
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc  Def  Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456 Def G89 Hi Jkl Mn12")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456 Def G89 Hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ". Abc~! Def# Ghi% Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.TitleCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc123 Def")
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc   Def   Ghi")
		})

		t.Run("convert ADa_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456 Def G 89 Hi Jkl Mn 12")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123-456 Def G 89 Hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ". Abc ~! Def # Ghi % Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.TitleCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "Abc 123 Def")
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = " "
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc  Def  Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "-"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456def G89hi Jkl Mn12")

			opts.Separators = "_"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456def G89hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~! Def# Ghi% Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.TitleCaseWithOptions("abc123def", opts)
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc _def _ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc -def -ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = " "
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc   Def   Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc _def _ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc -def -ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456def G 89hi Jkl Mn 12")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123-456def G 89hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc ~! Def # Ghi % Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc 456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = " "
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc  Def  Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc- Def- Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456 Def G89 Hi Jkl Mn12")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456 Def G89 Hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ". Abc~! Def# Ghi% Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert Title Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = " "
			result = stringcase.TitleCaseWithOptions("Abc Def Ghi", opts)
			assert.Equal(t, result, "Abc   Def   Ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts

			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc _ Def _ Ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc - Def - Ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123 456 Def G 89 Hi Jkl Mn 12")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc 123-456 Def G 89 Hi Jkl Mn 12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ". Abc ~! Def # Ghi % Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123 Abc 456 Def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc 456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
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
			result := stringcase.TitleCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "Abc Def Gh Ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert Ada_Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("Abc_Def_Ghi", opts)
			assert.Equal(t, result, "Abc_ Def_ Ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "_"
			result = stringcase.TitleCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "Abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc Def Ghi")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "Abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123 456def G89hi Jkl Mn12")

			opts.Keep = "-"
			result = stringcase.TitleCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "Abc123-456def G89hi Jkl Mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.TitleCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~! Def# Ghi% Jk Lm No ?")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.TitleCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123 Abc456 Def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.TitleCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
