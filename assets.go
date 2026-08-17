package assets

import (
	"embed"
	"net/http"
	"io/fs"
	"log"
)

//go:embed client/dist/*
var ClientFS embed.FS

func EmbedFS() http.Handler {
	distSubFS, err := fs.Sub(ClientFS, "client/dist")
	if err != nil {
		log.Fatalf("failed to create sub filesystem: %v", err)
	}

	fsHandler := http.FileServerFS(distSubFS)
	return fsHandler
}
