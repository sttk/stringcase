package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestCobolCase_convertCamelCase(t *testing.T) {
	result := stringcase.CobolCase("abcDefGHIjk")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCase_convertPascalCase(t *testing.T) {
	result := stringcase.CobolCase("AbcDefGHIjk")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCase_convertSnakeCase(t *testing.T) {
	result := stringcase.CobolCase("abc_def_ghi")
	assert.Equal(t, result, "ABC-DEF-GHI")
}

func TestCobolCase_convertKebabCase(t *testing.T) {
	result := stringcase.CobolCase("abc-def-ghi")
	assert.Equal(t, result, "ABC-DEF-GHI")
}

func TestCobolCase_convertTrainCase(t *testing.T) {
	result := stringcase.CobolCase("Abc-Def-Ghi")
	assert.Equal(t, result, "ABC-DEF-GHI")
}

func TestCobolCase_convertMacroCase(t *testing.T) {
	result := stringcase.CobolCase("ABC_DEF_GHI")
	assert.Equal(t, result, "ABC-DEF-GHI")
}

func TestCobolCase_convertCobolCase(t *testing.T) {
	result := stringcase.CobolCase("ABC-DEF-GHI")
	assert.Equal(t, result, "ABC-DEF-GHI")
}

func TestCobolCase_keepDigits(t *testing.T) {
	result := stringcase.CobolCase("abc123-456defG789HIJklMN12")
	assert.Equal(t, result, "ABC123-456DEF-G789-HI-JKL-MN12")
}

func TestCobolCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CobolCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "ABC-DEF-GHI-JK-LM-NO")
}

func TestCobolCase_convertEmpty(t *testing.T) {
	result := stringcase.CobolCase("")
	assert.Equal(t, result, "")
}

///

func TestCobolCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("abcDefGHIjk", "_-")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("AbcDefGHIjk", "_-")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "ABC_-DEF_-GHI")
}

func TestCobolCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "ABC_-DEF_-GHI")
}

func TestCobolCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.CobolCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.CobolCaseWithSep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "ABC123-456DEF-G789-HI-JKL-MN12")

	result = stringcase.CobolCaseWithSep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "ABC123--456DEF-G789-HI-JKL-MN12")
}

func TestCobolCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CobolCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, ".-ABC~!-DEF#-GHI%-JK-LM-NO-?")
}

func TestCobolCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.CobolCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestCobolCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("abcDefGHIjk", "_-")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("AbcDefGHIjk", "_-")
	assert.Equal(t, result, "ABC-DEF-GH-IJK")
}

func TestCobolCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "ABC_-DEF_-GHI")
}

func TestCobolCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "ABC_-DEF_-GHI")
}

func TestCobolCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "ABC-DEF-GHI")

	result = stringcase.CobolCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "ABC--DEF--GHI")
}

func TestCobolCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "ABC123-456DEF-G789-HI-JKL-MN12")

	result = stringcase.CobolCaseWithKeep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "ABC123--456DEF-G789-HI-JKL-MN12")
}

func TestCobolCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CobolCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, ".-ABC~!-DEF#-GHI%-JK-LM-NO-?")
}

func TestCobolCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.CobolCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}
