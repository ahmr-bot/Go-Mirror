package routers

import (
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	pkg.SetupCORS(&w)
	// http.ServeFile(w, r, "index.html")
	http.Error(w, "Welcome to MirrorEdge Mirror‘s API Server", 405)
}
