package kategorimenu

// KategoriMenu Struct
type KategoriMenu struct {
	ID           string   `json:"id"`
	CategoryName string   `json:"categoryname" validate:"nonzero"`
	ListMenu     []string `json:"listmenu"`
	Status       string   `json:"status"`
	Created      string   `json:"created"`
	Updated      string   `json:"updated"`
}
