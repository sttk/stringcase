package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestSnakeCaseWithNumsAsHead(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("abcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("AbcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("abc_def_ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("abc-def-ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert Trail-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("Abc-Def-Ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("ABC_DEF_GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("ABC-DEF-GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("123abc456def")
		assert.Equal(t, result, "123abc_456def")

		result = stringcase.SnakeCaseWithNumsAsHead("123ABC456DEF")
		assert.Equal(t, result, "123abc_456def")

		result = stringcase.SnakeCaseWithNumsAsHead("123Abc456Def")
		assert.Equal(t, result, "123_abc_456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsHead("")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNumsAsTail(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("abcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("AbcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("abc_def_ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("abc-def-ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert Trail-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("Abc-Def-Ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("ABC_DEF_GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("ABC-DEF-GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("123abc456def")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCaseWithNumsAsTail("123ABC456DEF")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCaseWithNumsAsTail("123Abc456Def")
		assert.Equal(t, result, "123_abc456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsTail("")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNumsAsWord(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("abcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("AbcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("abc_def_ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("abc-def-ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert Trail-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("Abc-Def-Ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("ABC_DEF_GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("ABC-DEF-GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("123abc456def")
		assert.Equal(t, result, "123_abc_456_def")

		result = stringcase.SnakeCaseWithNumsAsWord("123ABC456DEF")
		assert.Equal(t, result, "123_abc_456_def")

		result = stringcase.SnakeCaseWithNumsAsWord("123Abc456Def")
		assert.Equal(t, result, "123_abc_456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsWord("")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNumsAsPart(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("abcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("AbcDefGHIjk")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("abc_def_ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("abc-def-ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert Trail-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("Abc-Def-Ghi")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("ABC_DEF_GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("ABC-DEF-GHI")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("abc123-456defG89HIJklMN12")
		assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")
	})

	t.Run("convert with symbols as seperators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart(":.abc~!@def#$ghi%&jk(lm)no/?")
		assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("123abc456def")
		assert.Equal(t, result, "123abc456def")

		result = stringcase.SnakeCaseWithNumsAsPart("123ABC456DEF")
		assert.Equal(t, result, "123abc456def")

		result = stringcase.SnakeCaseWithNumsAsPart("123Abc456Def")
		assert.Equal(t, result, "123_abc456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNumsAsPart("")
		assert.Equal(t, result, "")
	})
}

///

func TestSnakeCaseWithNonAlphasAsHeadAndSep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_-def_-ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_-def_-ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc_123-456def_g_89hi_jkl_mn_12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep(":.abc~!@def#$ghi%&jk(lm)no/?",
			":@$&()/")
		assert.Equal(t, result, ".abc_~!_def_#_ghi_%_jk_lm_no_?")
	})

	t.Run("convert with starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("123abc456def", "-")
		assert.Equal(t, result, "123abc_456def")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("123ABC456DEF", "-")
		assert.Equal(t, result, "123abc_456def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndSep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsTailAndSep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc123-456_def_g89_hi_jkl_mn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep(":.abc~!@def#$ghi%&jk(lm)no/?",
			":@$&()/")
		assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
	})

	t.Run("convert with starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("123abc456def", "-")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndSep("123ABC456DEF", "-")
		assert.Equal(t, result, "123_abc456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndSep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsWordAndSep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc___def___ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc___def___ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc_123-456_def_g_89_hi_jkl_mn_12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep(":.abc~!@def#$ghi%&jk(lm)no/?",
			":@$&()/")
		assert.Equal(t, result, "._abc_~!_def_#_ghi_%_jk_lm_no_?")
	})

	t.Run("convert with starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("123abc456def", "-")
		assert.Equal(t, result, "123_abc_456_def")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndSep("123ABC456DEF", "-")
		assert.Equal(t, result, "123_abc_456_def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndSep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsPartAndSep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc123-456def_g89hi_jkl_mn12")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep(":.abc~!@def#$ghi%&jk(lm)no/?",
			":@$&()/")
		assert.Equal(t, result, ".abc~!_def#_ghi%_jk_lm_no_?")
	})

	t.Run("convert with starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("123abc456def", "-")
		assert.Equal(t, result, "123abc456def")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndSep("123ABC456DEF", "-")
		assert.Equal(t, result, "123abc456def")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndSep("", "-_")
		assert.Equal(t, result, "")
	})
}

///

func TestSnakeCaseWithNonAlphasAsHeadAndKeep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_-def_-ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_-def_-ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc_123_456def_g_89hi_jkl_mn_12")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc_123-456def_g_89hi_jkl_mn_12")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("123abc456def", "-")
		assert.Equal(t, result, "123abc_456def")

		result = stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("123ABC456DEF", "-")
		assert.Equal(t, result, "123abc_456def")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep(":.abc~!@def#$ghi%&jk(lm)no/?",
			".~!#%?")
		assert.Equal(t, result, ".abc_~!_def_#_ghi_%_jk_lm_no_?")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsHeadAndKeep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsTailAndKeep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc__def__ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc123_456_def_g89_hi_jkl_mn12")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc123-456_def_g89_hi_jkl_mn12")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("123abc456def", "-")
		assert.Equal(t, result, "123_abc456_def")

		result = stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("123ABC456DEF", "-")
		assert.Equal(t, result, "123_abc456_def")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep(":.abc~!@def#$ghi%&jk(lm)no/?",
			".~!#%?")
		assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsTailAndKeep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsWordAndKeep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc___def___ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc___def___ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc_-_def_-_ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc_123_456_def_g_89_hi_jkl_mn_12")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc_123-456_def_g_89_hi_jkl_mn_12")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("123abc456def", "-")
		assert.Equal(t, result, "123_abc_456_def")

		result = stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("123ABC456DEF", "-")
		assert.Equal(t, result, "123_abc_456_def")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep(":.abc~!@def#$ghi%&jk(lm)no/?",
			".~!#%?")
		assert.Equal(t, result, "._abc_~!_def_#_ghi_%_jk_lm_no_?")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsWordAndKeep("", "-_")
		assert.Equal(t, result, "")
	})
}

func TestSnakeCaseWithNonAlphasAsPartAndKeep(t *testing.T) {
	t.Run("convert camelCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert PascalCase", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("AbcDefGHIjk", "-_")
		assert.Equal(t, result, "abc_def_gh_ijk")
	})

	t.Run("convert snake_case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc_def_ghi", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc_def_ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert kebab-case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc-def-ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc-def-ghi", "-")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert Train-Case", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("Abc-Def-Ghi", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("Abc-Def-Ghi", "-")
		assert.Equal(t, result, "abc-_def-_ghi")
	})

	t.Run("convert MACRO_CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("ABC_DEF_GHI", "-")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("ABC_DEF_GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")
	})

	t.Run("convert COBOL-CASE", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("ABC-DEF-GHI", "_")
		assert.Equal(t, result, "abc_def_ghi")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("ABC-DEF-GHI", "-")
		assert.Equal(t, result, "abc-def-ghi")
	})

	t.Run("convert with keeping digits", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc123-456defG89HIJklMN12", "_")
		assert.Equal(t, result, "abc123_456def_g89hi_jkl_mn12")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("abc123-456defG89HIJklMN12", "-")
		assert.Equal(t, result, "abc123-456def_g89hi_jkl_mn12")
	})

	t.Run("convert when starting with digit", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("123abc456def", "-")
		assert.Equal(t, result, "123abc456def")

		result = stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("123ABC456DEF", "-")
		assert.Equal(t, result, "123abc456def")
	})

	t.Run("convert with symbols as separators", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep(":.abc~!@def#$ghi%&jk(lm)no/?",
			".~!#%?")
		assert.Equal(t, result, ".abc~!_def#_ghi%_jk_lm_no_?")
	})

	t.Run("convert an empty string", func(t *testing.T) {
		result := stringcase.SnakeCaseWithNonAlphasAsPartAndKeep("", "-_")
		assert.Equal(t, result, "")
	})
}
