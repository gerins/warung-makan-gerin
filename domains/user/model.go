package user

// User Struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
	Status   string `json:"status"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
