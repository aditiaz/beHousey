package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

		const MAX_UPLOAD_SIZE = 10 << 20

		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size is 10Mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")
		// get image filepath
		// dataContex := r.Context().Value("dataFile")
		// filepath := dataContex.(string)
		// Add your Cloudinary credentials ...
		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "uploads"})
		if err != nil {
			fmt.Println(err.Error())
		}

		// data := tempFile.Name()

		ctx = context.WithValue(r.Context(), "dataFile", resp.SecureURL)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
