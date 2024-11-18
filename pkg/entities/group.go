package entities

type Group struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type GroupUpdate struct {
	Name  *string `json:"name"`
	Users []User  `json:"users"`
}
