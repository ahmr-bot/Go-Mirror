package routers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	filePath := "root/" + mux.Vars(r)["path"]

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}
