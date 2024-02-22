package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestCamelCase_convertCamelCase(t *testing.T) {
	result := stringcase.CamelCase("abcDefGHIjk")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCase_convertPascalCase(t *testing.T) {
	result := stringcase.CamelCase("AbcDefGHIjk")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCase_convertSnakeCase(t *testing.T) {
	result := stringcase.CamelCase("abc_def_ghi")
	assert.Equal(t, result, "abcDefGhi")
}

func TestCamelCase_convertKebabCase(t *testing.T) {
	result := stringcase.CamelCase("abc-def-ghi")
	assert.Equal(t, result, "abcDefGhi")
}

func TestCamelCase_convertTrainCase(t *testing.T) {
	result := stringcase.CamelCase("Abc-Def-Ghi")
	assert.Equal(t, result, "abcDefGhi")
}

func TestCamelCase_convertMacroCase(t *testing.T) {
	result := stringcase.CamelCase("ABC_DEF_GHI")
	assert.Equal(t, result, "abcDefGhi")
}

func TestCamelCase_convertCobolCase(t *testing.T) {
	result := stringcase.CamelCase("ABC-DEF-GHI")
	assert.Equal(t, result, "abcDefGhi")
}

func TestCamelCase_keepDigits(t *testing.T) {
	result := stringcase.CamelCase("abc123-456defG789HIJklMN12")
	assert.Equal(t, result, "abc123456defG789HiJklMn12")
}

func TestCamelCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CamelCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "abcDefGhiJkLmNo")
}

func TestCamelCase_convertEmpty(t *testing.T) {
	result := stringcase.CamelCase("")
	assert.Equal(t, result, "")
}

///

func TestCamelCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("abcDefGHIjk", "_-")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("AbcDefGHIjk", "_-")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "abc_Def_Ghi")
}

func TestCamelCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abc_Def_Ghi")
}

func TestCamelCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.CamelCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.CamelCaseWithSep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "abc123-456defG789HiJklMn12")

	result = stringcase.CamelCaseWithSep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "abc123456defG789HiJklMn12")
}

func TestCamelCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CamelCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
}

func TestCamelCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.CamelCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestCamelCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("abcDefGHIjk", "_-")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("AbcDefGHIjk", "_-")
	assert.Equal(t, result, "abcDefGhIjk")
}

func TestCamelCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "abc_Def_Ghi")
}

func TestCamelCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "abc_Def_Ghi")
}

func TestCamelCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "abcDefGhi")

	result = stringcase.CamelCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "abc-Def-Ghi")
}

func TestCamelCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "abc123456defG789HiJklMn12")

	result = stringcase.CamelCaseWithKeep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "abc123-456defG789HiJklMn12")
}

func TestCamelCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.CamelCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
}

func TestCamelCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.CamelCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}
