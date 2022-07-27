package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dpull/coffer/filesystem"
	"golang.org/x/net/webdav"
)

type config struct {
	HttpAddr string
	Folder   string
	FSType   string
	FSParam  map[string]string
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	fd, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %+v", err)
	}
	defer fd.Close()
	log.SetOutput(io.MultiWriter(fd, os.Stdout))

	// TODO read from config
	config := config{
		HttpAddr: "localhost:8080",
		Folder:   "./temp",
		FSType:   "xor",
	}

	fs, err := filesystem.Create(config.FSType, config.Folder, config.FSParam)
	if err != nil {
		log.Fatalf("create file system failed: %+v", err)
	}

	err = http.ListenAndServe(config.HttpAddr, &webdav.Handler{
		FileSystem: fs,
		LockSystem: webdav.NewMemLS(),
	})
	if err != nil {
		log.Fatalf("http serve failed: %+v", err)
	}
}
