package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aneeshsunganahalli/Kumo"
	_ "modernc.org/sqlite"
)

type App struct {
	db *sql.DB
}

// itialize Database
func initDB() *sql.DB {
	connString := os.Getenv("DATABASE_URL")

	db ,err := sql.Open("sqlite", connString)
	if err != nil {
		log.Fatalf("Unable to open the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	log.Println("SUCCESS - Database Connected")	
	return db
}

func (a *App) root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SUCCESS -- API Running")
}

func main() {
	
	db := initDB()
	defer db.Close()

	app := &App{
		db: db,
	}


	// Embedding dist/ files into go binary
	fsHandler := assets.EmbedFS()

	http.Handle("/", fsHandler)
	http.HandleFunc("/root", app.root)

	log.Println("Server running at http://localhost:8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

