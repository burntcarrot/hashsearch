package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/burntcarrot/hashsearch/api/v1/handlers"
	"github.com/burntcarrot/hashsearch/api/v1/middleware"
	"github.com/burntcarrot/hashsearch/pkg/config"
	"github.com/burntcarrot/hashsearch/pkg/logging"
	imageRepository "github.com/burntcarrot/hashsearch/repository/image"
	"github.com/burntcarrot/hashsearch/usecase/image"
	"github.com/gorilla/mux"

	"github.com/tidwall/buntdb"
	"github.com/urfave/negroni"
)

func main() {
	// Get home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("failed to fetch home directory")
	}

	// Generate config path from the home directory.
	defaultConfigPath := path.Join(homeDir, ".hashsearch", "config.yml")

	// Define and parse flags.
	configPath := flag.String("config", defaultConfigPath, "Config file path")
	flag.Parse()

	// Load config.
	err = config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalln("failed to load config")
	}

	// Validate config.
	err = config.Validate()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize logger.
	logging.Logger = logging.InitLogger()

	// Open database connection.
	db, err := buntdb.Open(config.DB_URL)
	if err != nil {
		logging.Logger.Fatalln(err)
	}
	defer db.Close()

	// Create repository and service.
	imageRepo := imageRepository.NewBuntDB(db)
	imageService := image.NewService(imageRepo)

	// Create new router.
	r := mux.NewRouter()

	// Create a new Negroni instance.
	n := negroni.New(
		negroni.HandlerFunc(middleware.CORS),
		negroni.NewLogger(),
	)

	// Make handlers.
	handlers.MakeImageHandlers(r, *n, imageService)

	// Set router.
	http.Handle("/", r)

	// Set healthcheck route.
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create HTTP server with options.
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         config.SERVER_ADDR,
	}

	log.Printf("starting server on %s", config.SERVER_ADDR)

	// Start the server.
	err = server.ListenAndServe()
	if err != nil {
		logging.Logger.Fatalln(err)
	}
}
