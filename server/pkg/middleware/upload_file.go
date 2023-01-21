package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("image")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}

		defer file.Close()

		// 10MB SIZE
		const MAX_UPLOAD_SIZE = 10 << 20

		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		// buat file sementara untuk container file upload
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")

		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}

		defer tempFile.Close()

		// read all the content into a byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// tulis byte ke temp file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		// fileName := data[8:]

		// add filename to context

		ctx := context.WithValue(r.Context(), "dataFile", data)

		// lempar isi handle upload ke handler selanjutnya
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
