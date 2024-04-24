package model

type User struct {
	Name        string `gorm:"index:idx_name,unique" json:"name"`
	Password    string `gorm:"type:varchar(256)" json:"-"`
	Status      int8   `gorm:"type:tinyint" json:"status"`
	Role        int    `gorm:"type:tinyint" json:"role"`
	Email       string `gorm:"type:varchar(128)" json:"email"`
	Description string `gorm:"type:text" json:"description"`
	Extension   string `gorm:"type:text" json:"extension,omitempty"`
}

func (user *User) TableName() string { return "users" }
