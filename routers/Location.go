package routers

import (
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"net/http"
    "github.com/patrickmn/go-cache"
	"time"
	"encoding/json"
)
var (
	Cache = cache.New(5*time.Minute, 10*time.Minute)
)

var setlocation =  "xxx"
func HandleLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pkg.SetupCORS(&w)
    location := map[string]string{"location": setlocation}

    // Use a cache to store the location information
    cacheKey := "location"
    if location, found := Cache.Get(cacheKey); found {
    } else { 
		location = map[string]string{"location": setlocation}
		Cache.Set(cacheKey, location, cache.DefaultExpiration)
	}
    json.NewEncoder(w).Encode(location)
}