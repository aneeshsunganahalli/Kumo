package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aneeshsunganahalli/Kumo"
	"github.com/aneeshsunganahalli/Kumo/internal/database"
	"github.com/aneeshsunganahalli/Kumo/internal/handlers"
	_ "modernc.org/sqlite"
)

type App struct {
	db *db.DB
}

// initialize Database
func initDB() *db.DB {
	connString := os.Getenv("DATABASE_URL")

	sqlDB, err := sql.Open("sqlite", connString)
	if err != nil {
		log.Fatalf("Unable to open the database: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	log.Println("SUCCESS - Database Connected")
	return db.New(sqlDB)
}

func (a *App) root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SUCCESS -- API Running")
}

func main() {
	database := initDB()
	defer database.Close()

	app := &App{db: database}

	userH := &handlers.UserHandler{DB: database}
	prefH := &handlers.PreferenceHandler{DB: database}
	bgH := &handlers.BackgroundHandler{DB: database}
	audioH := &handlers.AudioHandler{DB: database}

	mux := http.NewServeMux()

	// User Handlers
	mux.HandleFunc("POST /api/users", userH.CreateUser)
	mux.HandleFunc("GET /api/users/{id}", userH.GetUser)
	mux.HandleFunc("PUT /api/users/{id}", userH.UpdateUser)

	// Preferences Handlers
	mux.HandleFunc("POST /api/users/{id}/preferences", prefH.CreatePreferences)
	mux.HandleFunc("GET /api/users/{id}/preferences", prefH.GetPreferences)
	mux.HandleFunc("PUT /api/users/{id}/preferences", prefH.UpdatePreferences)

	// Background Handlers
	mux.HandleFunc("POST /api/users/{id}/backgrounds", bgH.CreateBackground)
	mux.HandleFunc("GET /api/users/{id}/backgrounds", bgH.GetBackgrounds)
	mux.HandleFunc("GET /api/backgrounds/{id}", bgH.GetBackground)
	mux.HandleFunc("PUT /api/backgrounds/{id}", bgH.UpdateBackground)
	mux.HandleFunc("DELETE /api/backgrounds/{id}", bgH.DeleteBackground)

	// Audio Handlers
	mux.HandleFunc("POST /api/users/{id}/audio", audioH.CreateAudio)
	mux.HandleFunc("GET /api/users/{id}/audio", audioH.GetAudios)
	mux.HandleFunc("GET /api/audio/{id}", audioH.GetAudio)
	mux.HandleFunc("PUT /api/audio/{id}", audioH.UpdateAudio)
	mux.HandleFunc("DELETE /api/audio/{id}", audioH.DeleteAudio)

	// Embedding dist/ files into go binary
	fsHandler := assets.EmbedFS()

	mux.Handle("/", fsHandler)
	mux.HandleFunc("/root", app.root)

	log.Println("Server running at http://localhost:8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}
