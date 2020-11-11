package main

import (
	"GoPartice/errhandle/server/filelisting"
	"fmt"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error
type userError interface {
	error
	Message() string
}

func errorWraper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			code := http.StatusOK
			fmt.Sprintf("Error Handle request：%s", err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(w, userError.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc("/", errorWraper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
