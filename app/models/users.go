package models

type Users struct {
	Model

	FullName	string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Hash        string `json:"hash"`

	// FullName	string `gorm:"column:full_name;type:varchar(100)"  json:"full_name"`
	// Email       string `gorm:"type:varchar(100);unique_index" json:"email"`
	// Password    string `gorm:"type:varchar(50)" json:"password"`
	// Hash        string `gorm:"type:varchar(255);index:hash" json:"hash"`
}
