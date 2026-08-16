package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/aneeshsunganahalli/Kumo"
)


func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SUCCESS -- API Running")
}

func main() {
	http.HandleFunc("/root", root)

	// Embedding dist/ files into go binary
	distSubFS, err := fs.Sub(assets.ClientFS, "client/dist")
	if err != nil {
		log.Fatalf("failed to create sub filesystem: %v", err)
	}

	http.Handle("/", http.FileServerFS(distSubFS))

	log.Println("Server running at http://localhost:8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}