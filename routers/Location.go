package routers

import (
	"encoding/json"
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"log"
	"net/http"
)

func HandleLocation(w http.ResponseWriter, r *http.Request) {
	pkg.SetupCORS(&w)
	CacheConfig, err := redisClient.Get(redisClient.Context(), "config").Result()
	if err != nil {
		log.Fatal("从Redis读取配置文件缓存失败:", err)
	}
	var config Config
	if err := json.Unmarshal([]byte(CacheConfig), &config); err != nil {
		log.Fatal(err)
	}
	location := config.ServerLocation
	if location == "" {
		// Return an error if location information is not found
		http.Error(w, "Location is null", http.StatusInternalServerError)
		return
	}
	location = Location{
		ServerLocation: location,
	}
	json.NewEncoder(w).Encode(location)
}
