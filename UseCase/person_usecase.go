package usecase

import (
	"fmt"
	config "mereb/Config"
	Domain "mereb/Domain"
	Models "mereb/Domain/Models"

	"github.com/go-playground/validator/v10"
)

type PersonUsecase struct {
	personRepo Domain.PersonRepository
	infra      Domain.InfraStructure
}

func NewPersonUsecase(personRepo Domain.PersonRepository, infra Domain.InfraStructure) Domain.PersonUsecase {
	return &PersonUsecase{
		personRepo: personRepo,
		infra:      infra,
	}
}

func (usecase *PersonUsecase) GetAllPersons() ([]Models.Person, error) {
	return usecase.personRepo.GetAllPersons()

}
func (usecase *PersonUsecase) GetPersonByID(id string) (Models.Person, error) {
	return usecase.personRepo.GetPersonByID(id)
}

func (usecase *PersonUsecase) CreatePerson(person Models.Person) (Models.Person, error) {
	person.ID = usecase.infra.UUID()
	validationError := usecase.customErrorMessage(person)
	if validationError != nil {
		return Models.Person{}, validationError
	}
	_person, _err := usecase.personRepo.CreatePerson(person)
	return _person, _err
}

func (usecase *PersonUsecase) UpdatePerson(id string, newPerson Models.Person) (Models.Person, error) {
	_, err := usecase.personRepo.GetPersonByID(id)
	if err != nil {
		return Models.Person{}, err
	}
	newPerson.ID = id
	err = usecase.infra.ValidateStruct(newPerson)
	if err != nil {
		return Models.Person{}, err
	}
	_person, _err := usecase.personRepo.UpdatePerson(id, newPerson)
	return _person, _err
}
func (usecase *PersonUsecase) DeletePerson(id string) error {
	// check the existence of the user
	_, err := usecase.personRepo.GetPersonByID(id)
	if err != nil {
		return err
	}
	err = usecase.personRepo.DeletePerson(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *PersonUsecase) customErrorMessage(person Models.Person) error {

	err := usecase.infra.ValidateStruct(person)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var customErrors []string
			for _, err := range validationErrors {
				key := fmt.Sprintf("%s.%s", err.StructNamespace(), err.Tag())
				if msg, found := config.ErrorMessages[key]; found {
					customErrors = append(customErrors, msg)
				} else {
					customErrors = append(customErrors, err.Error())
				}
			}
			return fmt.Errorf("validation errors: %s", customErrors)
		}
	}
	return nil
}
