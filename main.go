package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/wao3/luogu-stats-card/handler"
)

//go:embed web/dist/*
var webFiles embed.FS

func main() {
	webFs, err := fs.Sub(webFiles, "web/dist")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/practice", handler.StatsHandler)
	http.HandleFunc("/api/guzhi", handler.GuzhiHandler)
	http.Handle("/", http.FileServer(http.FS(webFs)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}