package usecases

import (
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

func (pu *PersonUseCase) Register(newPerson *domain.NewPerson) (error) {
	err := pu.PersonRepository.Register(newPerson)
	return err
}


func (pu *PersonUseCase) GetAllPersons() (*[]domain.Person, error){
	persons, err := pu.PersonRepository.GetAllPersons()
	return persons, err
}

func (pu *PersonUseCase) GetPersonById(person_id string) (domain.Person, error) {
	foundPerson, err := pu.PersonRepository.GetPersonById(person_id)
	return foundPerson, err
}

func (pu *PersonUseCase) UpdatePerson(updatedInfo domain.NewPerson, person_id string) ([]domain.Person, error) {
	updatedPerson, err := pu.PersonRepository.UpdatePerson(updatedInfo, person_id)
	return updatedPerson, err
}

func (pu *PersonUseCase) DeletePerson(person_id string) (domain.Person, error) {
	deletedPerson, err := pu.PersonRepository.DeletePerson(person_id)
	return deletedPerson, err
}
