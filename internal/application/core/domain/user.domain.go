package domain

type User struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type UpdateUserDTO struct {
	Name    *string   `json:"name"`
	Age     *int      `json:"age"`
	Hobbies *[]string `json:"hobbies"`
}
