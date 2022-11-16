package handlers

import (
	dto "dumbsound/dto/result"
	usersdto "dumbsound/dto/user"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}
func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}
func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  users, err := h.UserRepository.FindUser()
  if err!= nil{
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(err.Error())
    return
  }
  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code : http.StatusOK,Data: users}
  json.NewEncoder(w).Encode(response)
}
func (h *handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-type","aplication/json")

request := new(usersdto.CreateUserRequest)
if err := json.NewDecoder(r.Body).Decode(&request)
err != nil{
  w.WriteHeader(http.StatusBadRequest)
  response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
  json.NewEncoder(w).Encode(response)
  return
}
validation := validator.New()
err := validation.Struct(request)
if err != nil {
  w.WriteHeader(http.StatusBadRequest)
  response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
  json.NewEncoder(w).Encode(response)
  return
}
user := models.User{
  Email: request.Email,
  Password: request.Password,
  FullName: request.FullName,
  Gender: request.Gender,
  Phone: request.Phone,
  Address: request.Address,
  Role: "user",
}
data,err := h.UserRepository.CreateUser(user)
if err != nil{
  w.WriteHeader(http.StatusInternalServerError)
  response := dto.ErrorResult{Code : http.StatusBadRequest,Message: err.Error()}
  json.NewEncoder(w).Encode(response)
  return
}
w.WriteHeader(http.StatusOK)
response := dto.SuccessResult{Code: http.StatusOK,Data: data}
json.NewEncoder(w).Encode(response)
}
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		FullName:  u.FullName,
		Email:    u.Email,
		Password: u.Password,
    Role : u.Role,
	}
  
}


