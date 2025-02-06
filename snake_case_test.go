package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestSnakeCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCase("abcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCase("AbcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCase("abc_def_ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCase("abc-def-ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert Trail-Case", func(t *testing.T) {
		result := stringcase.SnakeCase("Abc-Def-Ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCase("ABC_DEF_GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCase("ABC-DEF-GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.SnakeCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCase("123abc456def")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCase("123ABC456DEF")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCase("123Abc456Def")
		assert.Equal(t, result, "123_abc456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCase("")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert Trail-Case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.SnakeCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_abc_456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert Trail-Case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc456_def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc456_def")

			result = stringcase.SnakeCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_abc456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert Trail-Case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc_456_def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc_456_def")

			result = stringcase.SnakeCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_abc_456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert Trail-Case", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.SnakeCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123_abc456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.SnakeCaseWithOptions("", opts)
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123-456def_g_89hi_jkl_mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc_~!_def_#_ghi_%_jk_lm_no_?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.SnakeCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc_123def")
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456_def_g89_hi_jkl_mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc456_def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.SnakeCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123_def")
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc___def___ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc___def___ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123-456_def_g_89_hi_jkl_mn_12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._abc_~!_def_#_ghi_%_jk_lm_no_?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc_456_def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc_456_def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.SnakeCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc_123_def")
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "-"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")

			opts.Separators = "_"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def_g89hi_jkl_mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!_def#_ghi%_jk_lm_no_?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.SnakeCaseWithOptions("abc123def", opts)
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
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123-456def_g_89hi_jkl_mn_12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc_456def")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc_456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc_~!_def_#_ghi_%_jk_lm_no_?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc__def__ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456_def_g89_hi_jkl_mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc456_def")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc456_def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc___def___ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc___def___ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_-_def_-_ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc_123-456_def_g_89_hi_jkl_mn_12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123_abc_456_def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123_abc_456_def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "._abc_~!_def_#_ghi_%_jk_lm_no_?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
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
			result := stringcase.SnakeCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc_def_gh_ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "_"
			result = stringcase.SnakeCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")

			opts.Keep = "-"
			result = stringcase.SnakeCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def_g89hi_jkl_mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.SnakeCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.SnakeCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.SnakeCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!_def#_ghi%_jk_lm_no_?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.SnakeCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
