package main

import (
	"flag"
	"golang.org/x/net/webdav"
	"net/http"
)
/*
http.ListenAndServe(":8080", &webdav.Handler{
        FileSystem: webdav.Dir("."),
        LockSystem: webdav.NewMemLS(),
    })
*/
func main() {
	flag.Parse()
	fs := &webdav.Handler{
		FileSystem: webdav.Dir(""),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// username, password, ok := req.BasicAuth()

		if handleDirList(fs.FileSystem, w, req) {
			return
		}
		fs.ServeHTTP(w, req)
	})
	http.ListenAndServe(":8080", nil)

}
func handleDirList(fs webdav.FileSystem, w http.ResponseWriter, req *http.Request) bool {
	return false
}
