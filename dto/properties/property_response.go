package propertiesdto

import (
	"gorm.io/datatypes"
)


type PropertyResponse struct {	
	Name_Property string `json:"name_property"`
	City          string `json:"city" `
	Address       string `json:"address"`
	Price          float64     `json:"price"`
	TypeRent      string `json:"type_of_rent" `
	// Amenities     string `json:"amenities"`
	Amenities   datatypes.JSON `json:"amenities" `
	Bedroom     int    `json:"bedroom"`
	Bathroom    int    `json:"bathroom"`
	Sqf         string    `json:"sqf"`
	Description string `json:"description"`
}
