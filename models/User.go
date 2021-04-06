package models

type User struct {
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	Firstname  string `json:"Firstname"`
	Lastname   string `json:"Lastname"`
	Middlename string `json:"Middlename"`
	Email      string `json:"Email"`
	Phone      string `json:"Phone"`
}
