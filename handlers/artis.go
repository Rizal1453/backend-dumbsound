package handlers

import (
	artisdto "dumbsound/dto/artis"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerArtis struct {
	ArtisRepository repositories.ArtisRepository
}
func HandlerArtis(ArtisRepository repositories.ArtisRepository) *handlerArtis {
	return &handlerArtis{ArtisRepository}
}
func (h *handlerArtis) FindArtis(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Typr","application/json")

	artis,err := h.ArtisRepository.FindArtis()
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code : http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code : http.StatusOK,Data : artis}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerArtis)CreateArtis(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")


	request := new(artisdto.ArtisRequest)
	if err:= json.NewDecoder(r.Body).Decode(&request)
	err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	artis := models.Artis{
		Name : request.Name,
		Old: request.Old,
		Role:  request.Role,
		StartCareer: request.StartCareer,
	}
	artis,err = h.ArtisRepository.CreateArtis(artis)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code : http.StatusInternalServerError}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code :http.StatusOK,Data: artis}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerArtis)GetArtisById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplicatiom/json")

	id,_ :=strconv.Atoi(mux.Vars(r)["id"])

	
	artis,err := h.ArtisRepository.GetArtisById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response:= dto.ErrorResult{Code : http.StatusBadRequest, Message : err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response:= dto.SuccessResult{Code : http.StatusOK,Data : artis}
	json.NewEncoder(w).Encode(response)
	
}