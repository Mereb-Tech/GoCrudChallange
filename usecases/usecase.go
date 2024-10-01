package usecases

import (
	"errors"
	"mereb_go/domain"
)

func NewPersonUseCase(pr domain.PersonRepository) domain.PersonUseCase {
	return &PersonUseCase{
		PersonRepository: pr,
	}
}

type PersonUseCase struct {
	PersonRepository domain.PersonRepository
}

func (pu *PersonUseCase) DeletePerson(string) {
	panic("unimplemented")
}

func (pu *PersonUseCase) GetAllPersons() (*[]domain.Person, error){
	persons, err := pu.PersonRepository.GetAllPersons()
	return persons, err
}

func (pu *PersonUseCase) Register(*domain.NewPerson) (error) {
	panic("unimplemented")
}


func (pu *PersonUseCase) UpdatePerson(string) {
	panic("unimplemented")
}