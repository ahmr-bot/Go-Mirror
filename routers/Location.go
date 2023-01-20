package routers

import (
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"net/http"
    "github.com/patrickmn/go-cache"
	"time"
	"encoding/json"
	"os"
	"log"
)
var (
	Cache = cache.New(5*time.Minute, 10*time.Minute)
)
type Location struct {
	ServerLocation any `json:"server_location"`
}

var  config Config
func HandleLocation(w http.ResponseWriter, r *http.Request) {
	pkg.SetupCORS(&w)
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	// 关闭文件读取
	defer file.Close()
	// 解析 json 数据
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}
    location := config.ServerLocation
    if location == "" {
        // Return an error if location information is not found
        http.Error(w, "Location is null", http.StatusInternalServerError)
        return
    }
    // Use a cache to store the location information
    cacheKey := "location"
    if location, found := Cache.Get(cacheKey); found {
        location = location
    } else { 
		location = config.ServerLocation
		Cache.Set(cacheKey, location, cache.DefaultExpiration)
	} 		
	location = Location{
		ServerLocation: location,
	}
    json.NewEncoder(w).Encode(location)
}