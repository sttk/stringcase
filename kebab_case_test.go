package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestKebabCase_convertCamelCase(t *testing.T) {
	result := stringcase.KebabCase("abcDefGHIjk")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCase_convertPascalCase(t *testing.T) {
	result := stringcase.KebabCase("AbcDefGHIjk")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCase_convertSnakeCase(t *testing.T) {
	result := stringcase.KebabCase("abc_def_ghi")
	assert.Equal(t, result, "abc-def-ghi")
}

func TestKebabCase_convertKebabCase(t *testing.T) {
	result := stringcase.KebabCase("abc-def-ghi")
	assert.Equal(t, result, "abc-def-ghi")
}

func TestKebabCase_convertTrainCase(t *testing.T) {
	result := stringcase.KebabCase("Abc-Def-Ghi")
	assert.Equal(t, result, "abc-def-ghi")
}

func TestKebabCase_convertMacroCase(t *testing.T) {
	result := stringcase.KebabCase("ABC_DEF_GHI")
	assert.Equal(t, result, "abc-def-ghi")
}

func TestKebabCase_convertCobolCase(t *testing.T) {
	result := stringcase.KebabCase("ABC-DEF-GHI")
	assert.Equal(t, result, "abc-def-ghi")
}

func TestKebabCase_keepDigits(t *testing.T) {
	result := stringcase.KebabCase("abc123-456defG789HIJklMN12")
	assert.Equal(t, result, "abc123-456def-g789-hi-jkl-mn12")
}

func TestKebabCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.KebabCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "abc-def-ghi-jk-lm-no")
}

func TestKebabCase_convertEmpty(t *testing.T) {
	result := stringcase.KebabCase("")
	assert.Equal(t, result, "")
}

///

func TestKebabCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "abc_-def_-ghi")
}

func TestKebabCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abc_-def_-ghi")
}

func TestKebabCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.KebabCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.KebabCaseWithSep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "abc123-456def-g789-hi-jkl-mn12")

	result = stringcase.KebabCaseWithSep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "abc123--456def-g789-hi-jkl-mn12")
}

func TestKebabCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.KebabCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, ".-abc~!-def#-ghi%-jk-lm-no-?")
}

func TestKebabCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.KebabCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestKebabCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "abc-def-gh-ijk")
}

func TestKebabCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "abc_-def_-ghi")
}

func TestKebabCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abc_-def_-ghi")
}

func TestKebabCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abc-def-ghi")

	result = stringcase.KebabCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abc--def--ghi")
}

func TestKebabCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "abc123-456def-g789-hi-jkl-mn12")

	result = stringcase.KebabCaseWithKeep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "abc123--456def-g789-hi-jkl-mn12")
}

func TestKebabCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.KebabCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, ".-abc~!-def#-ghi%-jk-lm-no-?")
}

func TestKebabCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.KebabCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}
