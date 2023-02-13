package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	propertiesdto "housey/dto/properties"
	dto "housey/dto/result"
	"housey/models"
	"housey/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"gorm.io/datatypes"
)

type handlerProperty struct {
	PropertyRepository repositories.PropertyRepository
}

func HandlerProperty(PropertyRepository repositories.PropertyRepository) *handlerProperty {
	return &handlerProperty{PropertyRepository}
}

func (h *handlerProperty) FindProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	properties, err := h.PropertyRepository.FindProperties()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: properties}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) GetProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	property, err := h.PropertyRepository.GetProperty(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: property}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProperty) AddProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].( float64 ))
	// userRole := string(userInfo["listAs"].(string))

	price, _ := strconv.Atoi(r.FormValue("price"))
	bedroom, _ := strconv.Atoi(r.FormValue("bedroom"))
	bathroom, _ := strconv.Atoi(r.FormValue("bathroom"))

	request := propertiesdto.PropertyRequest{
		Name_Property: r.FormValue("name_property"),
		City:          r.FormValue("city"),
		Address:       r.FormValue("address"),
		Price:         float64(price),
		TypeRent:      r.FormValue("type_of_rent"),
		Amenities:     datatypes.JSON(r.FormValue("amenities")),
		Bedroom:       bedroom,
		Bathroom:      bathroom,
		Sqf:           r.FormValue("sqf"),
		Description:   r.FormValue("description"),
		// Image: filename,
		// User_Id: userId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	// get image filepath
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)
	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "uploads"})
	if err != nil {
		fmt.Println(err.Error())
	}

	property := models.Property{
		Name_Property: request.Name_Property,
		City:          request.City,
		Address:       request.Address,
		Price:         request.Price,
		TypeRent:      request.TypeRent,
		Amenities:     request.Amenities,
		Bedroom:       request.Bedroom,
		Bathroom:      request.Bathroom,
		Sqf:           request.Sqf,
		Description:   request.Description,
		Image:         resp.SecureURL,
		// User_Id: request.User_Id,

	}

	data, err := h.PropertyRepository.AddProperty(property)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.PropertyRepository.GetProperty(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
