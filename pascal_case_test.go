package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestPascalCase_convertCamelCase(t *testing.T) {
	result := stringcase.PascalCase("abcDefGHIjk")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCase_convertPascalCase(t *testing.T) {
	result := stringcase.PascalCase("AbcDefGHIjk")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCase_convertSnakeCase(t *testing.T) {
	result := stringcase.PascalCase("abc_def_ghi")
	assert.Equal(t, result, "AbcDefGhi")
}

func TestPascalCase_convertKebabCase(t *testing.T) {
	result := stringcase.PascalCase("abc-def-ghi")
	assert.Equal(t, result, "AbcDefGhi")
}

func TestPascalCase_convertTrainCase(t *testing.T) {
	result := stringcase.PascalCase("Abc-Def-Ghi")
	assert.Equal(t, result, "AbcDefGhi")
}

func TestPascalCase_convertMacroCase(t *testing.T) {
	result := stringcase.PascalCase("ABC_DEF_GHI")
	assert.Equal(t, result, "AbcDefGhi")
}

func TestPascalCase_convertCobolCase(t *testing.T) {
	result := stringcase.PascalCase("ABC-DEF-GHI")
	assert.Equal(t, result, "AbcDefGhi")
}

func TestPascalCase_keepDigits(t *testing.T) {
	result := stringcase.PascalCase("abc123-456defG789HIJklMN12")
	assert.Equal(t, result, "Abc123456DefG789HiJklMn12")
}

func TestPascalCase_convertWhenStartingWithDigit(t *testing.T) {
	result := stringcase.PascalCase("123abc456def")
	assert.Equal(t, result, "123Abc456Def")

	result = stringcase.PascalCase("123ABC456DEF")
	assert.Equal(t, result, "123Abc456Def")
}

func TestPascalCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.PascalCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "AbcDefGhiJkLmNo")
}

func TestPascalCase_convertEmpty(t *testing.T) {
	result := stringcase.PascalCase("")
	assert.Equal(t, result, "")
}

///

func TestPascalCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCaseWithSep_convertSnakeCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("abc_def_ghi", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithSep("abc_def_ghi", "-")
	assert.Equal(t, result, "Abc_Def_Ghi")
}

func TestPascalCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestPascalCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestPascalCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "Abc_Def_Ghi")
}

func TestPascalCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.PascalCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "Abc_Def_Ghi")
}

func TestPascalCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.PascalCaseWithSep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "Abc123456DefG789HiJklMn12")

	result = stringcase.PascalCaseWithSep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "Abc123-456DefG789HiJklMn12")
}

func TestPascalCaseWithSep_convertWhenStartingWithDigit(t *testing.T) {
	result := stringcase.PascalCaseWithSep("123abc456def", "_")
	assert.Equal(t, result, "123Abc456Def")

	result = stringcase.PascalCaseWithSep("123ABC456DEF", "_")
	assert.Equal(t, result, "123Abc456Def")
}

func TestPascalCaseWithSep_marksAsSeparators(t *testing.T) {
	result := stringcase.PascalCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
}

func TestPascalCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.PascalCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestPascalCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("AbcDefGHIjk", "-")
	assert.Equal(t, result, "AbcDefGhIjk")
}

func TestPascalCaseWithKeep_convertSnakeCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("abc_def_ghi", "-")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithKeep("abc_def_ghi", "_")
	assert.Equal(t, result, "Abc_Def_Ghi")
}

func TestPascalCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestPascalCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestPascalCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "Abc_Def_Ghi")
}

func TestPascalCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "AbcDefGhi")

	result = stringcase.PascalCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "Abc-Def-Ghi")
}

func TestPascalCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("abc123-456defG789HIJklMN12", "_")
	assert.Equal(t, result, "Abc123456DefG789HiJklMn12")

	result = stringcase.PascalCaseWithKeep("abc123-456defG789HIJklMN12", "-")
	assert.Equal(t, result, "Abc123-456DefG789HiJklMn12")
}

func TestPascalCaseWithKeep_convertWhenStartingWithDigit(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("123abc456def", "_")
	assert.Equal(t, result, "123Abc456Def")

	result = stringcase.PascalCaseWithKeep("123ABC456DEF", "_")
	assert.Equal(t, result, "123Abc456Def")
}

func TestPascalCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.PascalCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, ".Abc~!Def#Ghi%JkLmNo?")
}

func TestPascalCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.PascalCaseWithKeep("", "-_")
	assert.Equal(t, result, "")
}
