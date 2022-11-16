package handlers

import (
	authdto "dumbsound/dto/auth"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/pkg/bcrypt"
	jwtnih "dumbsound/pkg/jwt"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}
func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth)Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	request := new(authdto.RegisterRequest)
	if err:= json.NewDecoder(r.Body).Decode(&request)
	err != nil {w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code : http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	password,err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	user := models.User{
		Email: request.Email,
		Password: password,
		FullName : request.FullName,
		Gender : request.Gender,
		Phone : request.Phone,
		Address : request.Address,
		Role: "user",
	
	}
	data,err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
	
}
func (h *handlerAuth)Login(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")

request := new(authdto.LoginRequest)

if err := json.NewDecoder(r.Body).Decode(request)
err != nil {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Bad Request"}
	json.NewEncoder(w).Encode(response)
	return
}
user := models.User{
	Email: request.Email,
	Password: request.Password,
}
// check email
User,err := h.AuthRepository.Login(user.Email)
if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "email error"}
	json.NewEncoder(w).Encode(response)
	return
}
// check password
isValid := bcrypt.CheckPasswordHash(request.Password, User.Password)
if !isValid  {
	w.WriteHeader(http.StatusBadRequest)
	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "password salah"}
	json.NewEncoder(w).Encode(response)
	return
}
claims:= jwt.MapClaims{}
claims["id"] = User.ID
claims["exp"]= time.Now().Add(time.Hour * 2).Unix()

token , errGenerateToken:= jwtnih.GenerateToken(&claims)
if errGenerateToken != nil {
	log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
	return
}
loginResponse := authdto.LoginResponse{
	ID: User.ID,
	Email: User.Email,
	FullName: User.FullName,
	Password: User.Password,
	Token : token,
	Role: User.Role,
}
w.Header().Set("Content-Type", "application/json")
response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
json.NewEncoder(w).Encode(response)
}
func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.CheckAuth(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Gender:   user.Gender,
		Email:    user.Email,
		Phone:    user.Phone,
		Role: user.Role,
		Address:  user.Address,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}

