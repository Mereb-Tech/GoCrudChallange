package usecase

import (
	"fmt"
	Domain "mereb/Domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PersonUsecase struct {
	personRepo      Domain.PersonRepository
	personValidator *validator.Validate
}

func NewPersonUsecase(personRepo Domain.PersonRepository) Domain.PersonUsecase {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &PersonUsecase{
		personRepo:      personRepo,
		personValidator: validate,
	}

}

func (usecase *PersonUsecase) GetAllPersons() ([]Domain.Person, error) {
	return usecase.personRepo.GetAllPersons()

}
func (usecase *PersonUsecase) GetPersonByID(id string) (Domain.Person, error) {
	return usecase.personRepo.GetPersonByID(id)
}

func (usecase *PersonUsecase) CreatePerson(person Domain.Person) (Domain.Person, error) {
	person.ID = uuid.New().String()
	validationError := usecase.customErrorMessage(person)
	if validationError != nil {
		return Domain.Person{}, validationError
	}
	_person, _err := usecase.personRepo.CreatePerson(person)
	return _person, _err
}

func (usecase *PersonUsecase) UpdatePerson(id string, person Domain.Person) (Domain.Person, error) {
	person, err := usecase.personRepo.GetPersonByID(id)
	if err != nil {
		return Domain.Person{}, err
	}
	person.ID = id
	err = usecase.personValidator.Struct(person)
	if err != nil {
		return Domain.Person{}, err
	}
	_person, _err := usecase.personRepo.UpdatePerson(id, person)
	return _person, _err
}
func (usecase *PersonUsecase) DeletePerson(id string) error {
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

func (usecase *PersonUsecase) customErrorMessage(person Domain.Person) error {
	errorMessages := map[string]string{
		"Person.ID.required":   "User ID is required.",
		"Person.Name.required": "Name is required.",
		"Person.Name.min":      "Name must be at least 2 characters.",
		"Person.Name.max":      "Name can be a maximum of 100 characters.",
		"Person.Age.required":  "Age is required.",
		"Person.Age.gte":       "Age must be at least 1 year.",
		"Person.Age.lte":       "Age must be less than or equal to 150 years.",
		"Person.Hobbies.min":      "Each hobby must be a non-empty string.",
		"Person.Hobbies.max":      "Each hobby can be a maximum of 50 characters.",

	}

	err := usecase.personValidator.Struct(person)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var customErrors []string
			for _, err := range validationErrors {
				key := fmt.Sprintf("%s.%s", err.StructNamespace(), err.Tag())
				if msg, found := errorMessages[key]; found {
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
