package types

type User struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description,omitempty"`
}

type Password struct {
	OriginPassword  string `json:"origin_password"`
	CurrentPassword string `json:"current_password"`
}

type Role struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Roles struct {
	RoleIds []int64 `json:"role_ids"`
}

type Menu struct {
	Id          int64  `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Method      string `json:"method"`
	Description string `json:"description,omitempty"`
}

type Menus struct {
	MenuIds []int64 `json:"menu_ids"`
}
