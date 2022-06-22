package entity

import "github.com/google/uuid"

type User struct {
	UUID       uuid.UUID `json:"uuid" gorm:"primaryKey;not null"`
	Nama       string    `json:"nama" gorm:"not null;index"`
	Alamat     string    `json:"alamat" gorm:"not null"`
	Umur       int       `json:"umur" gorm:"not null"`
	Searchable string    `json:"searchable" gorm:"not null"`
}

type UserRequest struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Umur   int    `json:"umur"`
}
