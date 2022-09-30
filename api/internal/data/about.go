package data

type Imer struct {
	Name          string   `json:"name"`
	Lived         string   `json:"lived"`
	CurrentSchool string   `json:"current-school"`
	Study         string   `json:"study"`
	Email         string   `json:"email"`
	Hobbies       []string `json:"hobbies"`
	Interests     []string `json:"interests"`
}
