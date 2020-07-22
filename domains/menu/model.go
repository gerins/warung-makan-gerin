package menu

// Menu Struct
type Menu struct {
	ID       string `json:"id"`
	MenuName string `json:"menuname" validate:"nonzero"`
	Harga    int    `json:"harga" validate:"nonzero"`
	Stock    int    `json:"stock" validate:"nonzero"`
	Category string `json:"category"`
	Status   string `json:"status"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}
