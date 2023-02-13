package handlers

import (
	"encoding/json"
	"fmt"
	authdto "housey/dto/auth"
	dto "housey/dto/result"
	"housey/models"
	"housey/pkg/bcrypt"
	jwtToken "housey/pkg/jwt"
	"housey/repositories"
	"log"
	"net/http"
	"time"

	// "sync/atomic"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth)Register(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request);err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code:http.StatusBadRequest,Message:err.Error()}
		json.NewEncoder(w).Encode(response)	
		return
	}

	// hashing
	password,err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}


	user := models.User {
		Fullname: request.Fullname,
		Username: request.Username,
		Email: request.Email,
		Password: password,
		ListAs: request.ListAs,
		Gender:request.Gender,
		Phone: request.Phone,
		Address: request.Address,
		// Image: "default.png",
	}

	data,err := h.AuthRepository.Register(user)
	if err !=  nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	registerResponse := authdto.RegisterResponse{
		Username: data.Username,
		Message: "Sign Up Succesfully",
		ListAs: data.ListAs,
		
	}
w.WriteHeader(http.StatusOK)
response := dto.SuccessResult{Code:http.StatusOK,Data:registerResponse}
json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
       json.NewEncoder(w).Encode(response)
	   return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}
	// check username
	user,err := h.AuthRepository.Login(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code:http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// check password 
	isValid := bcrypt.CheckPasswordHash(request.Password,user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: "invalid username or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"]= user.ID
	claims["listAs"] = user.ListAs
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	
	token,errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		Username: user.Username,
		Token: token,
		ListAs: user.ListAs,
	}

 
	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}


