package repositories

import (
	"housey/models"
	"net/url"
	"strconv"

	"gorm.io/gorm"
)

type FilterRepository interface {
	MultiFilter(params url.Values) ([]models.Property, error)
}

func RepositoryFilter(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) MultiFilter(params url.Values) ([]models.Property, error) {
	var properties []models.Property

	type_of_rent := params.Get("type_of_rent")
	price, _ := strconv.ParseFloat(params.Get("price"), 64)
	bedroom, _ := strconv.Atoi(params.Get("bedroom"))
	bathroom, _ := strconv.Atoi(params.Get("bathroom"))
	amenities := params.Get("amenities")

	err := r.db.Where("type_rent = ? AND price < ? AND bedroom = ? AND bathroom = ? AND amenities = ?", type_of_rent, price, bedroom, bathroom, amenities).Find(&properties).Error

	return properties, err

}
