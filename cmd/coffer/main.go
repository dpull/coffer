package main

import (
	"encoding/json"
	"flag"
	"github.com/dpull/coffer/filesystem"
	"golang.org/x/net/webdav"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type config struct {
	HttpAddr string            `json:"http_addr"`
	Folder   string            `json:"folder"`
	FSType   string            `json:"fs_type"`
	FSParam  map[string]string `json:"fs_param"`
}

func initLog() (*os.File, error) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	fd, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	log.SetOutput(io.MultiWriter(fd, os.Stdout))
	return fd, nil
}

func readConfig(conf *config) error {
	path := flag.String("c", "config.json", "config file")
	flag.Parse()

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, conf)
}

func main() {
	fd, err := initLog()
	if err != nil {
		log.Fatalf("init log failed:%+v", err)
	}
	defer fd.Close()

	var conf config
	err = readConfig(&conf)
	if err != nil {
		flag.Usage()
		log.Fatalf("read config failed:%=v", err)
	}

	fs, err := filesystem.Create(conf.FSType, conf.Folder, conf.FSParam)
	if err != nil {
		log.Fatalf("create file system failed: %+v", err)
	}

	err = http.ListenAndServe(conf.HttpAddr, &webdav.Handler{
		FileSystem: fs,
		LockSystem: webdav.NewMemLS(),
	})
	if err != nil {
		log.Fatalf("http serve failed: %+v", err)
	}
}
