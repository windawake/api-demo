package Models

type Profile struct {
	ID     int    `gorm:"primary_key" json:"id"`
	UserId string `json:"user_id"`
	CityId int    `json:"city_id"`
	Hobby  string `json:"hobby"`
	Age    string `json:"age"`
}
