package repository

import (
	models "GoCrudChallenge/Domain/Models"
	"fmt"
	"log"
)

type PersonRepository struct {
	person_persistence []*models.Person
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		person_persistence: []*models.Person{},
	}
}

func (persistence *PersonRepository) CreatePerson(person *models.Person) error {
	persistence.person_persistence = append(persistence.person_persistence, person)
	return nil
}

func (persistence *PersonRepository) GetPersonByID(id string) (*models.Person, error) {
	for _, person := range persistence.person_persistence {
		if person.ID == id {
			return person, nil
		}
	}
	return nil, fmt.Errorf("Person not found")
}

func (persistence *PersonRepository) UpdatePerson(person *models.Person) error {

	for i, existing_person := range persistence.person_persistence {
		if existing_person.ID == person.ID {
			log.Println("existing", existing_person, "new", person)
			
			if person.Name != "" {
				existing_person.Name = person.Name
			}
			if person.Age != 0 {
				existing_person.Age = person.Age
			}
			if person.Hobbies != nil {
				existing_person.Hobbies = person.Hobbies
			}
			persistence.person_persistence[i] = existing_person
			return nil
		}
	}
	return fmt.Errorf("Person not found")
}

func (persistence *PersonRepository) DeletePerson(id string) error {
	for i, person := range persistence.person_persistence {
		if person.ID == id {
			persistence.person_persistence = append(persistence.person_persistence[:i], persistence.person_persistence[i+1:]...)
			return nil
		}
	}
	return nil
}

func (persistence *PersonRepository) GetAllPersons() []*models.Person {
	return persistence.person_persistence
}
