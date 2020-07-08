package user

// User Struct
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"nonzero, min=4, max=16"`
	Password string `json:"password" validate:"nonzero, min=4, max=16"`
	Status   string `json:"status"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type UserToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
