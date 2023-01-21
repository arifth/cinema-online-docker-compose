package handlers

import (
	"context"
	"encoding/json"
	filmdto "final-task/dto/film"
	dto "final-task/dto/result"
	"final-task/models"
	"final-task/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Filmhandler struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *Filmhandler {
	return &Filmhandler{FilmRepository}
}

func (h *Filmhandler) FindFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	film, err := h.FilmRepository.FindFilm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	// for i, p := range film {
	// 	imagePath := os.Getenv("PATH_FILE") + p.Image
	// 	film[i].Image = imagePath
	// }
	if err != nil {
		fmt.Println("errornya adalah", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *Filmhandler) FindSingleFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.FindSingleFilm(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	film.Image = os.Getenv("PATH_FILE") + film.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: film,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Filmhandler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfp["id"].(float64))

	dataContext := r.Context().Value("dataFile")

	// assign nama file ke variable filename
	filepath := dataContext.(string)

	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetour"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// get data country convrt ke int
	dataPrice, _ := strconv.Atoi(r.FormValue("price"))
	dataCategory_id, _ := strconv.Atoi(r.FormValue("category_id"))

	request := filmdto.CreateFilmRequest{
		Title:       r.FormValue("title"),
		Price:       dataPrice,
		Description: r.FormValue("description"),
		// Image:       "",
		CategoryId: dataCategory_id,
		FilmUrl:    r.FormValue("film_url"),
		Thumbnail:  r.FormValue("thumbnail"),
	}

	// validate request against struct form created
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	film := models.Film{
		Title:      request.Title,
		Price:      request.Price,
		Image:      resp.SecureURL,
		CategoryId: dataCategory_id,
		// Category:    models.CategoryResponse{},
		FilmUrl:     request.FilmUrl,
		Description: request.Description,
		Thumbnail:   "default.jpg",
	}

	data, err := h.FilmRepository.CreateFilm(film)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}

// /* Handler for request ,logic => get all values from req body ,decode it , put it in user variable which it's datatype is user model,
// then write it to DB with UpdateUser() method
// then return response to user with succesCode and data written with NewEncoder().Encode()
// */

// // COMMENT: able to insert Name ,but Email and password isn't included
// // solved , caused by typo in dto.SuccessResult

// FIXED: create new value in database,instead updating it
func (h *Filmhandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContext := r.Context().Value("dataFile")

	// assign nama file ke variable filename
	filename := dataContext.(string)

	// NOTE: face error caused by key value in postman using whitespace after it , DONT DO THAT !!

	// get data country convert ke int
	dataPrice, _ := strconv.Atoi(r.FormValue("price"))

	request := filmdto.CreateFilmRequest{
		Title:       r.FormValue("title"),
		Price:       dataPrice,
		Description: r.FormValue("description"),
		Image:       filename,
	}

	// NOTES: lakukan pengecekan isian form dengan if
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film := models.Film{}

	// check all field for emptieness
	if request.Title != "" {
		film.Title = request.Title
	}
	if request.Price != 0 {
		film.Price = request.Price
	}

	if request.Description != "" {
		film.Description = request.Description
	}
	if request.Image != "" {
		film.Image = request.Image
	}

	// fmt.Println(request)

	data, err := h.FilmRepository.UpdateFilm(film, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())

	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *Filmhandler) DeleteFilm(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.FindSingleFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FilmRepository.DeleteFilm(film, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
