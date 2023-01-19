package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"os"
	_ "path/filepath"
	_ "strings"
	"time"
)

type Config struct {
	Directories []struct {
		Path        string `json:"path"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"directories"`
	ServerAddr any `json:"server_addr"`
}
type Dir struct {
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Path        string   `json:"path"`
	Files       []Files  `json:"files"`
	Directories []string `json:"directories"`
}

var (
	dirCache = cache.New(5*time.Minute, 10*time.Minute)
)

type Files struct {
	Name string
	Url  string
	Size int64
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	// 读取配置文件
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
	dirCache.Set("config", config, cache.DefaultExpiration)
	if x, found := dirCache.Get("config"); found {
		config = x.(Config)
	} else {
	}
	dirPath := "root/" + mux.Vars(r)["path"]
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		http.Error(w, "Directory does not exist", http.StatusNotFound)
		return
	}

	var dir Dir
	var description string
	var image string
	for _, d := range config.Directories {
		if d.Path == dirPath {
			description = d.Description
			image = d.Image
			break
		}
	}
	if x, found := dirCache.Get(dirPath); found {
		dir = x.(Dir)
	} else {
		var directories []string
		currentDir, err := os.Open(dirPath)
		if err != nil {
			log.Fatal(err)
		}
		defer currentDir.Close()
		fileInfos, err := currentDir.Readdir(-1)
		if err != nil {
			log.Fatal(err)
		}
		var files []Files
		for _, fi := range fileInfos {
			if fi.IsDir() {
				directories = append(directories, fi.Name())
			} else {
				if mux.Vars(r)["path"] == "" {
					files = append(files, Files{
						Name: fi.Name(),
						Url:  fmt.Sprintf("%s/download%s/%s", config.ServerAddr, mux.Vars(r)["path"], fi.Name()),
						Size: fi.Size() / 1024 / 1024,
					})
				} else {
					files = append(files, Files{
						Name: fi.Name(),
						Url:  fmt.Sprintf("%s/download/%s/%s", config.ServerAddr, mux.Vars(r)["path"], fi.Name()),
						Size: fi.Size() / 1024 / 1024,
					})
				}
			}
		}
		dir = Dir{
			Description: description,
			Image:       image,
			Path:        dirPath,
			Files:       files,
			Directories: directories,
		}
		dirCache.Set(dirPath, dir, cache.DefaultExpiration)
	}
	json.NewEncoder(w).Encode(dir)
}
