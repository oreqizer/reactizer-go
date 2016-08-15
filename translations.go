package main

import (
  "path/filepath"
  "os"
  "flag"
	"io/ioutil"
	"regexp"
	"strings"
	"encoding/json"
	"sort"

	"github.com/golang/glog"
)

type Message struct {
	Id string						`json:"id"`
	Translation string	`json:"translation"`
}

type Messages []Message

func (m Messages) Len() int {
	return len(m)
}

func (m Messages) Less(i, j int) bool {
	return strings.Compare(m[i].Id, m[j].Id) < 0
}

func (m Messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
  flag.Parse()
	messages := Messages{}
	// visit function
	visit := func (path string, f os.FileInfo, err error) error {
		// only walk files named 'errors.go'
		if f.IsDir() || strings.HasPrefix(path, ".") || !strings.HasSuffix(path, "/errors.go") {
			return nil
		}
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		// take all quoted values
		r := regexp.MustCompile(`".*"`)
		translations := r.FindAll(b, -1)
		for _, t := range translations {
			// remove redundant quotes
			id := strings.Replace(string(t), "\"", "", -1)
			// add to messages
			messages = append(messages, Message{Id: id})
		}
		return nil
	}

  err := filepath.Walk(".", visit)
	if err != nil {
		glog.Error(err)
		return
	}

	// sort it by id
	sort.Sort(messages)

	// marshall & write JSON to file
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		glog.Error(err)
		return
	}
	err = ioutil.WriteFile("./locales/_raw.untranslated.json", data, os.ModePerm)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}
