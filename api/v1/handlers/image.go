package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"github.com/burntcarrot/hashsearch/api/v1/presenter"
	"github.com/burntcarrot/hashsearch/entity"
	"github.com/burntcarrot/hashsearch/pkg/config"
	"github.com/burntcarrot/hashsearch/pkg/logging"
	"github.com/burntcarrot/hashsearch/usecase/image"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var (
	ErrListImages = "failed to list images"
	ErrFindImages = "failed to find images"
	ErrGetData    = "failed to get data"
	ErrReadInput  = "failed to read input"
)

// search is the HTTP handler for serving an image search request.
func search(service image.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		// Parse multipart form.
		err = r.ParseMultipartForm(1024)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Get file from the form.
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Generate file path.
		filePath := filepath.Join(config.FILES_DIR, handler.Filename)

		// Create the new file.
		targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		// Copy the file data to the new file.
		if _, err := io.Copy(targetFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create image entry using image service.
		err = service.Create(filePath)
		if err != nil {
			logging.Logger.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(ErrListImages))
			return
		}

		// Get distances with respect to the current image.
		data, err := service.GetDistances(filePath)
		if err != nil {
			logging.Logger.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(ErrListImages))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(ErrFindImages))
			return
		}

		// Convert returned data into a response using the presenter.
		var present []*presenter.Image
		for _, d := range data {
			present = append(present, &presenter.Image{
				Path:     d.Path,
				Hash:     fmt.Sprintf("%064b", d.Hash),
				Distance: d.Distance,
			})
		}

		// Sort response with respect to the image distance.
		sort.Slice(present, func(i, j int) bool {
			return present[i].Distance < present[j].Distance
		})

		// Set headers.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		// Generate JSON response using the presenter.
		if err := json.NewEncoder(w).Encode(present); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(ErrGetData))
		}
	})
}

// list is the HTTP handler for listing all images.
func list(service image.Usecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.Image
		var err error

		// Get list of images.
		data, err = service.List()
		if err != nil {
			logging.Logger.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(ErrListImages))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(ErrFindImages))
			return
		}

		// Convert returned data into a response using the presenter.
		var present []*presenter.Image
		for _, d := range data {
			present = append(present, &presenter.Image{
				Path: d.Path,
				Hash: fmt.Sprintf("%064b", d.Hash),
			})
		}

		// Set headers.
		w.Header().Set("Content-Type", "application/json")

		// Generate JSON response using the presenter.
		if err := json.NewEncoder(w).Encode(present); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(ErrGetData))
		}
	})
}

func MakeImageHandlers(r *mux.Router, n negroni.Negroni, service image.Usecase) {
	// Create a subrouter.
	v1 := r.PathPrefix("/v1").Subrouter()

	// Register handlers

	v1.Handle("/list", n.With(
		negroni.Wrap(list(service)),
	)).Methods("GET", "OPTIONS").Name("list")

	v1.Handle("/search", n.With(
		negroni.Wrap(search(service)),
	)).Methods("POST", "OPTIONS").Name("search")
}
