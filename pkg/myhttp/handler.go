package myhttp

import (
	"fmt"
	"hsr/loto/internal/logger"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type MyHttp struct {
	Path string
}

func GetHttpHandler() *MyHttp {
	return &MyHttp{}
}

func (s MyHttp) GetHandler(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	fmt.Println("Test GetHandler")

	path, err := filepath.Abs(s.Path)
	if err != nil {
		logger.Error.Print(err)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error.Print(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
	_, err = w.Write(file)
	if err != nil {
		logger.Error.Print(err)
	}

	return w
}
