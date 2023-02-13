package models

import (
	"gorm.io/datatypes"
)

type Property struct {
	ID            int            `json:"id"`
	Name_Property string         `json:"name_property"`
	City          string         `json:"city" `
	Address       string         `json:"address"`
	Price         float64        `json:"price"`
	TypeRent      string         `json:"type_of_rent" `
	Amenities     datatypes.JSON `json:"amenities" `
	Bedroom       int            `json:"bedroom"`
	Bathroom      int            `json:"bathroom"`
	Sqf           string         `json:"sqf"`
	Description   string         `json:"description"`
	Image         string         `json:"image" `
}

func (Property) TableName() string {
	return "properties"
}
