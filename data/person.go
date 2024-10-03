package data

// Person represents a person with an ID, Name, Age, and Hobbies
type Person struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Age     int      `json:"age"`
    Hobbies []string `json:"hobbies"`
}

// Dummy data: a map of persons 
var Persons = map[string]Person{
    "1": {
        ID:      "1",
        Name:    "Bisrat Berhanu",
        Age:     30,
        Hobbies: []string{"Reading", "Traveling"},
    },
    "2": {
        ID:      "2",
        Name:    "Melake Berhanu",
        Age:     25,
        Hobbies: []string{"Cooking", "Running"},
    },
}
