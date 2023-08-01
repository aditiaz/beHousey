package middleware

import (
	"context"
	"fmt"
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
			http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
			return
		}
		defer file.Close()

		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "uploads"})
		if err != nil {
			fmt.Println(err.Error())
		}

		ctx = context.WithValue(r.Context(), "dataFile", resp.SecureURL)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
