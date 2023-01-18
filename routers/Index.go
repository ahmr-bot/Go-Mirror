package routers

import "net/http"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "index.html")
	http.Error(w, "Welcome to MirrorEdge Mirror‘s API Server", 405)
}
