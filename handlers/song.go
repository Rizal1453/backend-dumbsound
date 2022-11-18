package handlers

import (
	"context"
	dto "dumbsound/dto/result"
	songdto "dumbsound/dto/song"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"github.com/go-playground/validator/v10"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
)

	

type handlerSong struct {
	SongRepository repositories.SongRepository
}
func HandlerSong(SongRepository repositories.SongRepository) *handlerSong{
	return &handlerSong{SongRepository}
}
func (h *handlerSong)FindSong(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")

song,err:= h.SongRepository.FindSong()
if err != nil{
	w.WriteHeader(http.StatusInternalServerError)
	response := dto.ErrorResult{Code :http.StatusBadRequest,Message: err.Error()}
	json.NewEncoder(w).Encode(response)
	return
}
for i, p := range song {
		imagePath := os.Getenv("PATH_FILE") + p.Image
		song[i].Image = imagePath
	}
	for i, p := range song {
		imagePath := os.Getenv("PATH_FILE") + p.Song
		song[i].Song = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: song}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerSong)CreateSong(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	
	dataThumbnail := r.Context().Value("dataFile") // add this code
	filepath := dataThumbnail.(string)            // add thImageode
	
	dataSong := r.Context().Value("songFile") // add this code
	fileSong := dataSong.(string)             // add this code
	
	
	artisId, _ := strconv.Atoi(r.FormValue("artis_id"))
	
	
	request := songdto.SongRequest{
		Title:    r.FormValue("title"),
		Year:     r.FormValue("year"),

		ArtisID: artisId,
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dumbSound"});
	if err != nil {
 	 fmt.Println(err.Error())
	}

	song := models.Song{
		Title:    request.Title,
		Image:    resp.SecureURL,
		Year:     request.Year,
		Song:     fileSong,
		ArtisID : request.ArtisID,
	}

	data, err := h.SongRepository.CreateSong(song)
	if err != nil {
		// w.Header().Set("Content-type", "aplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
func (h *handlerSong) GetSongById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	song, err := h.SongRepository.GetSongById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	song.Image = os.Getenv("PATH_FILE") + song.Image
	song.Song = os.Getenv("PATH_FILE") + song.Song

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: song}
	json.NewEncoder(w).Encode(response)
}
