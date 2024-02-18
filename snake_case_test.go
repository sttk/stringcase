package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestSnakeCase_convertCamelCase(t *testing.T) {
	result := stringcase.SnakeCase("abcDefGHIjk")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCase_convertPascalCase(t *testing.T) {
	result := stringcase.SnakeCase("AbcDefGHIjk")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCase_convertSnakeCase(t *testing.T) {
	result := stringcase.SnakeCase("abc_def_ghi")
	assert.Equal(t, result, "abc_def_ghi")
}

func TestSnakeCase_convertKebabCase(t *testing.T) {
	result := stringcase.SnakeCase("abc-def-ghi")
	assert.Equal(t, result, "abc_def_ghi")
}

func TestSnakeCase_convertTrainCase(t *testing.T) {
	result := stringcase.SnakeCase("Abc-Def-Ghi")
	assert.Equal(t, result, "abc_def_ghi")
}

func TestSnakeCase_convertMacroCase(t *testing.T) {
	result := stringcase.SnakeCase("ABC_DEF_GHI")
	assert.Equal(t, result, "abc_def_ghi")
}

func TestSnakeCase_convertCobolCase(t *testing.T) {
	result := stringcase.SnakeCase("ABC-DEF-GHI")
	assert.Equal(t, result, "abc_def_ghi")
}

func TestSnakeCase_convertKeepDigits(t *testing.T) {
	result := stringcase.SnakeCase("abc123-456defG89HIJklMN12")
	assert.Equal(t, result, "abc123_456def_g89_hi_jkl_mn12")
}

func TestSnakeCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.SnakeCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "abc_def_ghi_jk_lm_no")
}

func TestSnakeCase_convertEmpty(t *testing.T) {
	result := stringcase.SnakeCase("")
	assert.Equal(t, result, "")
}

///

func TestSnakeCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "abc__def__ghi")
}

func TestSnakeCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abc__def__ghi")
}

func TestSnakeCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithSep_convertKeepDigits(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("abc123-456defG89HIJklMN12", "-")
	assert.Equal(t, result, "abc123_456def_g89_hi_jkl_mn12")

	result = stringcase.SnakeCaseWithSep("abc123-456defG89HIJklMN12", "_")
	assert.Equal(t, result, "abc123-_456def_g89_hi_jkl_mn12")
}

func TestSnakeCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.SnakeCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
}

func TestSnakeCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.SnakeCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestSnakeCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "abc_def_gh_ijk")
}

func TestSnakeCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "abc__def__ghi")
}

func TestSnakeCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abc__def__ghi")
}

func TestSnakeCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abc_def_ghi")

	result = stringcase.SnakeCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abc-_def-_ghi")
}

func TestSnakeCaseWithKeep_convertKeepDigits(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("abc123-456defG89HIJklMN12", "_")
	assert.Equal(t, result, "abc123_456def_g89_hi_jkl_mn12")

	result = stringcase.SnakeCaseWithKeep("abc123-456defG89HIJklMN12", "-")
	assert.Equal(t, result, "abc123-_456def_g89_hi_jkl_mn12")
}

func TestSnakeCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, "._abc~!_def#_ghi%_jk_lm_no_?")
}

func TestSnakeCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.SnakeCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}
