package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmr-bot/MirrorsAPI/middleware"
	"github.com/ahmr-bot/MirrorsAPI/routers"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	ListenPort  string `json:"listen_port"`
	Directories []struct {
		Path        string `json:"path"`
		Description string `json:"description"`
	} `json:"directories"`
}

// 定义缓存
var (
	dirCache = cache.New(5*time.Minute, 10*time.Minute)
)

// Start~
func main() {
	// 读取配置文件
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
		fmt.Errorf("配置文件读取失败")
	}
	fmt.Printf("配置文件加载成功")
	// 关闭文件读取
	defer file.Close()
	// 解析 json 数据
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}
	dirCache.Set("config", config, cache.DefaultExpiration)
	// 对配置文件进行缓存
	fmt.Println("配置文件写入缓存")
	fmt.Println("\n镜缘镜像站 API 服务启动成功！")
	// 读取缓存
	if x, found := dirCache.Get("config"); found {
		config = x.(Config)
	} else {
		log.Fatal("无法从缓存中找到配置文件")
	}
	fmt.Printf("监听端口" + config.ListenPort)
	// 设定路由
	router := mux.NewRouter()
	router.HandleFunc("/", middleware.RateLimiter(60, 60)(routers.HandleIndex))
	router.HandleFunc("/list/{path:.*}", middleware.RateLimiter(60, 60)(routers.HandleList))
	router.HandleFunc("/download/{path:.*}", middleware.RateLimiter(5, 60)(routers.HandleDownload))
	http.ListenAndServe(":"+config.ListenPort, router)
}
