package test

import (
	"mereb/Domain"
	Mocks "mereb/Test/Mocks"
	usecase "mereb/UseCase"

	"github.com/stretchr/testify/suite"
)

type usecaseSuite struct {
	suite.Suite
	usecase Domain.PersonUsecase
}

func (suite *usecaseSuite) SetupTest() {
	repository := new(Mocks.PersonRepository)
	suite.usecase = usecase.NewPersonUsecase(repository,)
}

func (suite *usecaseSuite)  TestCreatePersonPositive()