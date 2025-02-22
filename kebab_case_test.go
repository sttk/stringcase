package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestKebabCase(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.KebabCase("abcDefGHIjk")
		assert.Equal(t, result, "abc-def-gh-ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.KebabCase("AbcDefGHIjk")
		assert.Equal(t, result, "abc-def-gh-ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.KebabCase("abc_def_ghi")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.KebabCase("abc-def-ghi")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.KebabCase("Abc-Def-Ghi")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.KebabCase("ABC_DEF_GHI")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.KebabCase("ABC-DEF-GHI")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.KebabCase("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.KebabCase(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.KebabCase("123abc456def")
		assert.Equal(t, result, "123-abc456-def")

		result = stringcase.KebabCase("123ABC456DEF")
		assert.Equal(t, result, "123-abc456-def")

		result = stringcase.KebabCase("123Abc456Def")
		assert.Equal(t, result, "123-abc456-def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.KebabCase("")
		assert.Equal(t, result, "")
	})
}

func TestKebabCaseWithOptions(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc-456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc-456def")

			result = stringcase.KebabCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-abc-456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc456-def")

			result = stringcase.KebabCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-abc456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc-456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc-456-def")

			result = stringcase.KebabCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-abc-456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.KebabCaseWithOptions("123Abc456Def", opts)
			assert.Equal(t, result, "123-abc456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc-456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc-456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc-123def")
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc123-def")
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc-456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc-456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc-123-def")
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc-456def")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc-456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc456-def")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc456-def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc-456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc-456-def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
			assert.Equal(t, result, "abc-123-def")
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "-"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")

			opts.Separators = "_"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.KebabCaseWithOptions("abc123def", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_def-_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456def-g-89hi-jkl-mn-12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc-456def")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc-456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_-def_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456-def-g89-hi-jkl-mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc456-def")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc456-def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-_-def-_-ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc---def---ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc-123-456-def-g-89-hi-jkl-mn-12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123-abc-456-def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123-abc-456-def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".-abc-~!-def-#-ghi-%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
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
			result := stringcase.KebabCaseWithOptions("abcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("AbcDefGHIjk", opts)
			assert.Equal(t, result, "abc-def-gh-ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("abc_def_ghi", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc-def-ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("Abc-Def-Ghi", opts)
			assert.Equal(t, result, "abc--def--ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "_"
			result = stringcase.KebabCaseWithOptions("ABC_DEF_GHI", opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("ABC-DEF-GHI", opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")

			opts.Keep = "-"
			result = stringcase.KebabCaseWithOptions("abc123-456defG89HIJklMN12", opts)
			assert.Equal(t, result, "abc123-456def-g89hi-jkl-mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.KebabCaseWithOptions("123abc456def", opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.KebabCaseWithOptions("123ABC456DEF", opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.KebabCaseWithOptions(":.abc~!@def#$ghi%&jk(lm)no/?", opts)
			assert.Equal(t, result, ".abc~!-def#-ghi%-jk-lm-no-?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.KebabCaseWithOptions("", opts)
			assert.Equal(t, result, "")
		})
	})
}
