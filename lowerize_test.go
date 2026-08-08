package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestLowerize(t *testing.T) {
	t.Run("non-alphabets as head of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456def.g.89hi.jkl.mn.12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "abc.def.ghi.jk.lm.no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc.456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc.456def")

			result = stringcase.Lowerize("123Abc456Def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as tail of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456.def.g89.hi.jkl.mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "abc.def.ghi.jk.lm.no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc456.def")

			result = stringcase.Lowerize("123Abc456Def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as a word", func(t *testing.T) {
		t.Skip()
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: true,
			SeparateAfterNonAlphabets:  true,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456.def.g.89.hi.jkl.mn.12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "abc.def.ghi.jk.lm.no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")

			result = stringcase.Lowerize("123Abc456Def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})
	})

	t.Run("non-alphabets as part of a word", func(t *testing.T) {
		opts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456def.g89hi.jkl.mn12")
		})

		t.Run("convert with symbols as seperators", func(t *testing.T) {
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "abc.def.ghi.jk.lm.no")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.Lowerize("123Abc456Def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456def.g.89hi.jkl.mn.12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456def.g.89hi.jkl.mn.12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc.456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc.456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
			assert.Equal(t, result, "abc.123def")
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456.def.g89.hi.jkl.mn12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456.def.g89.hi.jkl.mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
			assert.Equal(t, result, "abc123.def")
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456.def.g.89.hi.jkl.mn.12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456.def.g.89.hi.jkl.mn.12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
			assert.Equal(t, result, "abc.123.def")
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456def.g89hi.jkl.mn12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456def.g89hi.jkl.mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			t.Skip()
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456def.g.89hi.jkl.mn.12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456def.g.89hi.jkl.mn.12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			t.Skip()
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc.456def")

			opts.Keep = "-"
			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc.456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			t.Skip()
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			t.Skip()
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456.def.g89.hi.jkl.mn12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456.def.g89.hi.jkl.mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")

			opts.Keep = "_"
			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc456.def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456.def.g.89.hi.jkl.mn.12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456.def.g.89.hi.jkl.mn.12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
			assert.Equal(t, result, "abc.123.def")
		})
	})

	t.Run("non-alphabets as part of a word and with separators", func(t *testing.T) {
		t.Skip()
		origOpts := stringcase.Options{
			SeparateBeforeNonAlphabets: false,
			SeparateAfterNonAlphabets:  false,
		}

		t.Run("convert camelCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "_"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "-"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Separators = "_"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456def.g89hi.jkl.mn12")

			opts.Separators = "_"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456def.g89hi.jkl.mn12")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Separators = ":@$&()/"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert with starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})

		t.Run("alphabets and numbers in separators are no effect", func(t *testing.T) {
			opts := origOpts
			opts.Separators = "-b2"
			result := stringcase.Lowerize("abc123def", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._def._ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-def.-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456def.g.89hi.jkl.mn.12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456def.g.89hi.jkl.mn.12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc.456def")

			opts.Keep = "-"
			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc.456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc_.def_.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456.def.g89.hi.jkl.mn12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456.def.g89.hi.jkl.mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc456.def")

			opts.Keep = "_"
			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc456.def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc._.def._.ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.-.def.-.ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123.456.def.g.89.hi.jkl.mn.12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc.123-456.def.g.89.hi.jkl.mn.12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123.abc.456.def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, "..abc.~!.def.#.ghi.%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
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
			result := stringcase.Lowerize("abcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert PascalCase", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("AbcDefGHIjk", '.', opts)
			assert.Equal(t, result, "abc.def.gh.ijk")
		})

		t.Run("convert snake_case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("abc_def_ghi", '.', opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert kebab-case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc-def-ghi", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert Train-Case", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("Abc-Def-Ghi", '.', opts)
			assert.Equal(t, result, "abc-.def-.ghi")
		})

		t.Run("convert MACRO_CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "_"
			result = stringcase.Lowerize("ABC_DEF_GHI", '.', opts)
			assert.Equal(t, result, "abc_def_ghi")
		})

		t.Run("convert COBOL-CASE", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc.def.ghi")

			opts.Keep = "-"
			result = stringcase.Lowerize("ABC-DEF-GHI", '.', opts)
			assert.Equal(t, result, "abc-def-ghi")
		})

		t.Run("convert with keeping digits", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "_"
			result := stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123.456def.g89hi.jkl.mn12")

			opts.Keep = "-"
			result = stringcase.Lowerize("abc123-456defG89HIJklMN12", '.', opts)
			assert.Equal(t, result, "abc123-456def.g89hi.jkl.mn12")
		})

		t.Run("convert when starting with digit", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-"
			result := stringcase.Lowerize("123abc456def", '.', opts)
			assert.Equal(t, result, "123abc456def")

			result = stringcase.Lowerize("123ABC456DEF", '.', opts)
			assert.Equal(t, result, "123abc456def")
		})

		t.Run("convert with symbols as separators", func(t *testing.T) {
			opts := origOpts
			opts.Keep = ".~!#%?"
			result := stringcase.Lowerize(":.abc~!@def#$ghi%&jk(lm)no/?", '.', opts)
			assert.Equal(t, result, ".abc~!.def#.ghi%.jk.lm.no.?")
		})

		t.Run("convert an empty string", func(t *testing.T) {
			opts := origOpts
			opts.Keep = "-_"
			result := stringcase.Lowerize("", '.', opts)
			assert.Equal(t, result, "")
		})
	})
}
