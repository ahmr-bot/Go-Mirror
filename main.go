package main

import (
	"encoding/json"
	"github.com/ahmr-bot/MirrorsAPI/middleware"
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"github.com/ahmr-bot/MirrorsAPI/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Config struct {
	ListenPort  string `json:"listen_port"`
	Directories []struct {
		Path        string `json:"path"`
		Description string `json:"description"`
	} `json:"directories"`
	RedisAddr      string      `json:"redis_addr"`
	ServerAddr     any         `json:"server_addr"`
	RedisPass      string      `json:"redis_pass"`
	RedisDB        int         `json:"redis_db"`
	location       any         `json:"server_location"`
	ServerLocation interface{} `json:"server_location"`
}

func (c Config) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// 使其他文件也可以访问redisClient
var redisClient = pkg.RedisClient()

func init() {
	log.Println("初始化中...")
	//  读取配置文件
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		log.Fatal("配置文件读取失败")
	}
	log.Println("配置文件读取成功")
	defer file.Close()
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}
	// 检查redis是否连接成功
	_, err = redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		log.Fatal("Redis连接失败:", err)
	} else {
		log.Println("Redis连接成功")
	}
	// 将配置文件存入redis
	_, err = redisClient.Set(redisClient.Context(), "config", config, 0).Result()
	if err != nil {
		log.Fatal("配置文件存入Redis失败:", err)
	} else {
		log.Println("配置文件存入Redis成功")
	}
}

// Start~
func main() {
	// 从redis中读取config
	CacheConfig, err := redisClient.Get(redisClient.Context(), "config").Result()
	if err != nil {
		log.Fatal("从Redis读取配置文件失败:", err)
	} else {
		log.Println("从Redis读取配置文件成功")
	}
	var config Config
	if err := json.Unmarshal([]byte(CacheConfig), &config); err != nil {
		log.Fatal(err)
	}
	log.Printf("监听端口" + config.ListenPort)
	// 设定路由
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.RateLimiter(60, 60)(routers.HandleIndex))
	router.HandleFunc("/api/location", middleware.RateLimiter(60, 60)(routers.HandleLocation))
	router.HandleFunc("/api/list/{path:.*}", middleware.RateLimiter(60, 60)(routers.HandleList))
	router.HandleFunc("/api/download/{path:.*}", middleware.RateLimiter(5, 60)(routers.HandleDownload))
	// http.ListenAndServe(":"+config.ListenPort, router)
	// listen and serve on then print
	log.Println("API 服务已启动")
	log.Fatal(http.ListenAndServeTLS(":"+config.ListenPort, "cert.pem", "cert.key", router))

}
