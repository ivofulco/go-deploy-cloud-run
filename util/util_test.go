package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type utilSuite struct {
	suite.Suite
}

func (suite *utilSuite) TestIsValidCEP() {
	cep := "01153000"
	isValid := IsValidCEP(cep)
	assert.True(suite.T(), isValid)
}

func (suite *utilSuite) TestIsInValidCEP() {
	cep := "2921607a"
	isValid := IsValidCEP(cep)
	assert.False(suite.T(), isValid)

	cep = "a921607"
	isValid = IsValidCEP(cep)
	assert.False(suite.T(), isValid)

	cep = ""
	isValid = IsValidCEP(cep)
	assert.False(suite.T(), isValid)
}

func (suite *utilSuite) TestIsDigit() {
	cep := "00000000"
	isValid := IsDigit(cep)
	assert.True(suite.T(), isValid)
}

func (suite *utilSuite) TestIsNotDigit() {
	cep := "133aa350"
	isValid := IsDigit(cep)
	assert.False(suite.T(), isValid)
}

func (suite *utilSuite) TestGetEnvVar() {
	varName := "PORT"
	value := GetEnvVariable(varName)
	assert.NotEmpty(suite.T(), value)

	varName = "XPTO"
	value = GetEnvVariable(varName)
	assert.Empty(suite.T(), value)
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(utilSuite))
}
