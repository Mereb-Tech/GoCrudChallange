package Domain

type Person struct {
	ID      string   `json:"uid" validate:"required"`
	Name    string   `json:"name" validate:"required,min=2,max=100"`
	Age     int      `json:"age" validate:"gte=0,lte=150"`
	Hobbies []string `json:"hobbies" validate:"dive,required,min=1,max=50"`
}

type PersonRepository interface {
	GetAllPerson() ([]Person, error)
	GetPersonByID(id string) (Person, error)
	CreatePerson(person Person) (Person, error)
	UpdatePerson(id string, person Person) (Person, error)
	DeletePerson(id string) (Person, error)
}

type PersonUsecase interface {
	GetAllPerson() ([]Person, error)
	GetPersonByID(id string) (Person, error)
	CreatePerson(person Person) (Person, error)
	UpdatePerson(id string, person Person) (Person, error)
	DeletePerson(id string) (Person, error)
}
