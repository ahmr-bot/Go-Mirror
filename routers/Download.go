package routers

import (
	"github.com/ahmr-bot/MirrorsAPI/pkg"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	pkg.SetupCORS(&w)
	filePath := "root/" + mux.Vars(r)["path"]
	// 1. 设置响应头
	w.Header().Set("Content-Type", "application/octet-stream")
	// 2. 设置 Content-Disposition 响应头，告诉浏览器下载文件
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	// 3. 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	// 4. 将文件内容复制到响应中
	io.Copy(w, f)
}
