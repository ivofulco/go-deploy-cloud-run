package cep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type viaCepSuite struct {
	suite.Suite
	viaCep ViaCep
}

func (suite *viaCepSuite) SetupSuite() {
	suite.viaCep = *InstanceViaCep()
}

func (suite *viaCepSuite) TestGetCepInfo() {
	cep := "01153000"
	location, err := suite.viaCep.FindLocation(cep)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "São Paulo,SP", location)
}

func (suite *viaCepSuite) TestNotGetCepInfo() {
	cep := "98765432"
	_, err := suite.viaCep.FindLocation(cep)
	assert.Error(suite.T(), err)

	cep = "aaa"
	_, err = suite.viaCep.FindLocation(cep)
	assert.Error(suite.T(), err)
}

func TestViaCepSuite(t *testing.T) {
	suite.Run(t, new(viaCepSuite))
}
