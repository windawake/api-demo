package Models

type Role struct {
	ID      int    `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	IsSuper int    `json:"is_super"`
}
