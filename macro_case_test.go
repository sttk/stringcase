package stringcase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sttk/stringcase"
)

func TestMacroCase_convertCamelCase(t *testing.T) {
	result := stringcase.MacroCase("abcDefGHIjk")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCase_convertPascalCase(t *testing.T) {
	result := stringcase.MacroCase("AbcDefGHIjk")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCase_convertSnakeCase(t *testing.T) {
	result := stringcase.MacroCase("abc_def_ghi")
	assert.Equal(t, result, "ABC_DEF_GHI")
}

func TestMacroCase_convertKebabCase(t *testing.T) {
	result := stringcase.MacroCase("abc-def-ghi")
	assert.Equal(t, result, "ABC_DEF_GHI")
}

func TestMacroCase_convertTrainCase(t *testing.T) {
	result := stringcase.MacroCase("Abc-Def-Ghi")
	assert.Equal(t, result, "ABC_DEF_GHI")
}

func TestMacroCase_convertMacroCase(t *testing.T) {
	result := stringcase.MacroCase("ABC_DEF_GHI")
	assert.Equal(t, result, "ABC_DEF_GHI")
}

func TestMacroCase_convertCobolCase(t *testing.T) {
	result := stringcase.MacroCase("ABC-DEF-GHI")
	assert.Equal(t, result, "ABC_DEF_GHI")
}

func TestMacroCase_keepDigits(t *testing.T) {
	result := stringcase.MacroCase("abc123-456defG89HIJklMN12")
	assert.Equal(t, result, "ABC123_456DEF_G89_HI_JKL_MN12")
}

func TestMacroCase_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.MacroCase(":.abc~!@def#$ghi%&jk(lm)no/?")
	assert.Equal(t, result, "ABC_DEF_GHI_JK_LM_NO")
}

func TestMacroCase_convertEmpty(t *testing.T) {
	result := stringcase.MacroCase("")
	assert.Equal(t, result, "")
}

///

func TestMacroCaseWithSep_convertCamelCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCaseWithSep_convertPascalCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCaseWithSep_convertKebabCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("abc-def-ghi", "-")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithSep("abc-def-ghi", "_")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithSep_convertTrainCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithSep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithSep_convertMacroCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithSep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "ABC__DEF__GHI")
}

func TestMacroCaseWithSep_convertCobolCase(t *testing.T) {
	result := stringcase.MacroCaseWithSep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithSep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithSep_keepDigits(t *testing.T) {
	result := stringcase.MacroCaseWithSep("abc123-456defG89HIJklMN12", "-")
	assert.Equal(t, result, "ABC123_456DEF_G89_HI_JKL_MN12")

	result = stringcase.MacroCaseWithSep("abc123-456defG89HIJklMN12", "_")
	assert.Equal(t, result, "ABC123-_456DEF_G89_HI_JKL_MN12")
}

func TestMacroCaseWithSep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.MacroCaseWithSep(":.abc~!@def#$ghi%&jk(lm)no/?", ":@$&()/")
	assert.Equal(t, result, "._ABC~!_DEF#_GHI%_JK_LM_NO_?")
}

func TestMacroCaseWithSep_convertEmpty(t *testing.T) {
	result := stringcase.MacroCaseWithSep("", "-_")
	assert.Equal(t, result, "")
}

///

func TestMacroCaseWithKeep_convertCamelCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("abcDefGHIjk", "-_")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCaseWithKeep_convertPascalCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("AbcDefGHIjk", "-_")
	assert.Equal(t, result, "ABC_DEF_GH_IJK")
}

func TestMacroCaseWithKeep_convertKebabCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("abc-def-ghi", "_")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithKeep("abc-def-ghi", "-")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithKeep_convertTrainCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("Abc-Def-Ghi", "_")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithKeep("Abc-Def-Ghi", "-")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithKeep_convertMacroCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("ABC_DEF_GHI", "-")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithKeep("ABC_DEF_GHI", "_")
	assert.Equal(t, result, "ABC__DEF__GHI")
}

func TestMacroCaseWithKeep_convertCobolCase(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("ABC-DEF-GHI", "_")
	assert.Equal(t, result, "ABC_DEF_GHI")

	result = stringcase.MacroCaseWithKeep("ABC-DEF-GHI", "-")
	assert.Equal(t, result, "ABC-_DEF-_GHI")
}

func TestMacroCaseWithKeep_keepDigits(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("abc123-456defG89HIJklMN12", "_")
	assert.Equal(t, result, "ABC123_456DEF_G89_HI_JKL_MN12")

	result = stringcase.MacroCaseWithKeep("abc123-456defG89HIJklMN12", "-")
	assert.Equal(t, result, "ABC123-_456DEF_G89_HI_JKL_MN12")
}

func TestMacroCaseWithKeep_treatMarksAsSeparators(t *testing.T) {
	result := stringcase.MacroCaseWithKeep(":.abc~!@def#$ghi%&jk(lm)no/?", ".~!#%?")
	assert.Equal(t, result, "._ABC~!_DEF#_GHI%_JK_LM_NO_?")
}

func TestMacroCaseWithKeep_convertEmpty(t *testing.T) {
	result := stringcase.MacroCaseWithKeep("", "_-")
	assert.Equal(t, result, "")
}
