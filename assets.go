package assets

import "embed"

//go:embed client/dist/*
var ClientFS embed.FS
