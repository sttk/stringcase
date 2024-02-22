package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestTrainCase_convertCamelCase(t *testing.T) {
	result := stringcase.TrainCase("abcDefGHIjk")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCase_convertPascalCase(t *testing.T) {
	result := stringcase.TrainCase("AbcDefGHIjk")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCase_convertSnakeCase(t *testing.T) {
	result := stringcase.TrainCase("abc_def_ghi")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestTrainCase_convertKebabCase(t *testing.T) {
	result := stringcase.TrainCase("abc-def-ghi")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestTrainCase_convertTrainCase(t *testing.T) {
	result := stringcase.TrainCase("Abc-Def-Ghi")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestTrainCase_convertMacroCase(t *testing.T) {
	result := stringcase.TrainCase("ABC_DEF_GHI")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestTrainCase_convertCobolCase(t *testing.T) {
	result := stringcase.TrainCase("ABC-DEF-GHI")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestTrainCase_keepDigits(t *testing.T) {
	result := stringcase.TrainCase("abc123-456defG89HIJklMN12")
	assert.Equal(t, result, "Abc123-456def-G89-Hi-Jkl-Mn12")
}

func TestTrainCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.TrainCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "Abc-Def-Ghi-Jk-Lm-No")
}

func TestTrainCase_convertEmpty(t *testing.T) {
	result := stringcase.TrainCase("")
	assert.Equal(t, result, "")
}

func TestTrainCase_startsWithDigit(t *testing.T) {
	result := stringcase.TrainCase("123Abc_456Def")
	assert.Equal(t, result, "123-Abc-456-Def")
}

///

func TestTrainCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "Abc_-Def_-Ghi")
}

func TestTrainCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "Abc_-Def_-Ghi")
}

func TestTrainCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.TrainCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.TrainCaseWithSep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "Abc123-456def-G789-Hi-Jkl-Mn12")

	result = stringcase.TrainCaseWithSep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "Abc123--456def-G789-Hi-Jkl-Mn12")
}

func TestTrainCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.TrainCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, ".-Abc~!-Def#-Ghi%-Jk-Lm-No?")
}

func TestTrainCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.TrainCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

func TestTrainCaseWithSep_startsWithDigit(t *testing.T) {
	result := stringcase.TrainCaseWithSep("123Abc_456Def", "_")
	assert.Equal(t, result, "123-Abc-456-Def")

	result = stringcase.TrainCaseWithSep("123Abc_456Def", "-")
	assert.Equal(t, result, "123-Abc_-456-Def")
}

///

func TestTrainCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "Abc-Def-Gh-Ijk")
}

func TestTrainCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "Abc_-Def_-Ghi")
}

func TestTrainCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "Abc_-Def_-Ghi")
}

func TestTrainCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")

	result = stringcase.TrainCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "Abc--Def--Ghi")
}

func TestTrainCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "Abc123-456def-G789-Hi-Jkl-Mn12")

	result = stringcase.TrainCaseWithKeep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "Abc123--456def-G789-Hi-Jkl-Mn12")
}

func TestTrainCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.TrainCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, ".-Abc~!-Def#-Ghi%-Jk-Lm-No?")
}

func TestTrainCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}

func TestTrainCaseWithKeep_startsWithDigit(t *testing.T) {
	result := stringcase.TrainCaseWithKeep("123Abc_456Def", "-")
	assert.Equal(t, result, "123-Abc-456-Def")

	result = stringcase.TrainCaseWithKeep("123Abc_456Def", "_")
	assert.Equal(t, result, "123-Abc_-456-Def")
}
