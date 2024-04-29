package types

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type Password struct {
	OriginPassword  string `json:"origin_password"`
	CurrentPassword string `json:"current_password"`
}

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Id   int64  `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}
