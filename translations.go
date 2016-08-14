package main

import (
  "path/filepath"
  "os"
  "flag"
	"io/ioutil"
	"regexp"
	"strings"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
)

type message struct {
	id string
	translation string
}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() || strings.HasPrefix(path, ".") || !strings.HasSuffix(path, "errors.go") {
		return nil
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	r := regexp.MustCompile(`".*"`)
	translations := r.FindAll(b, -1)
	messages := []message{}
	for _, t := range translations {
		messages = append(messages, message{id: t})
	}
	data, err := json.Marshal(messages) // TODO indent
	if err != nil {
		return err
	}
	fmt.Print(data)
	return nil
}

func main() {
  flag.Parse()
  err := filepath.Walk(".", visit)
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Info("Success")
	return
}
