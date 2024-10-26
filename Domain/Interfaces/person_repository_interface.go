package interfaces

import models "GoCrudChallenge/Domain/Models"


type PersonRepository interface {
    CreatePerson(person *models.Person) error
    GetPersonByID(id string) (*models.Person, error)
    UpdatePerson(person *models.Person) error
    DeletePerson(id string) error
    GetAllPersons() []*models.Person
}
