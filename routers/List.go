package routers

import (
	"encoding/json"
	"fmt"
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	_ "path/filepath"
	_ "strings"
)

func HandleList(w http.ResponseWriter, r *http.Request) {

	pkg.SetupCORS(&w)
	CacheConfig, err := redisClient.Get(redisClient.Context(), "config").Result()
	if err != nil {
		log.Fatal("从Redis读取配置文件失败:", err)
	}
	var config Config
	if err := json.Unmarshal([]byte(CacheConfig), &config); err != nil {
		log.Fatal(err)
	}
	dirPath := "root/" + mux.Vars(r)["path"]
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		http.Error(w, "目录不存在", http.StatusNotFound)
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
	// 如果 redis 中有 dirPath 的，则直接从redis中读取 若没有，则从文件中读取 并将读取的结果存入redis
	CacheDir, err := redisClient.Get(redisClient.Context(), dirPath).Result()
	if err != nil {
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
						Url:  fmt.Sprintf("%s:%s/api/download%s/%s", config.ServerAddr, config.ListenPort, mux.Vars(r)["path"], fi.Name()),
						Size: fi.Size() / 1024 / 1024,
					})
				} else {
					files = append(files, Files{
						Name: fi.Name(),
						Url:  fmt.Sprintf("%s:%s/api/download/%s/%s", config.ServerAddr, config.ListenPort, mux.Vars(r)["path"], fi.Name()),
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
		// 将dir存入redis
		redisClient.Set(redisClient.Context(), dirPath, dir, 0)
		json.NewEncoder(w).Encode(dir)
	} else {
		json.NewEncoder(w).Encode(CacheDir)
	}
}
