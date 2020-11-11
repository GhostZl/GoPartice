package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type fileError string

func (err fileError) Message() string {
	return string(err)
}
func (err fileError) Error() string {
	return err.Message()
}
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) < 0 {
		return fileError("路径错误，请以list开头")
	}
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
