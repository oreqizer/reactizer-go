package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Message struct {
	Id          string `json:"id"`
	Translation string `json:"translation"`
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

func fillMessages(m *Messages, files string) func(string, os.FileInfo, error) error {
	return func(path string, f os.FileInfo, err error) error {
		// only walk files named 'errors.go'
		if f.IsDir() || !strings.HasSuffix(path, "/"+files) {
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
			*m = append(*m, Message{Id: id})
		}
		return nil
	}
}

func main() {
	flag.Parse()
	output := flag.String("output", "message_ids.json", "Output file")
	files := flag.String("files", "errors.go", "Files to search")
	log.SetOutput(os.Stderr)

	// search files
	messages := &Messages{}
	err := filepath.Walk(".", fillMessages(messages, *files))
	if err != nil {
		log.Print(err)
		return
	}

	// sort it by id
	sort.Sort(messages)

	// marshall & write JSON to file
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		log.Print(err)
		return
	}
	err = ioutil.WriteFile(*output, data, os.ModePerm)
	if err != nil {
		log.Print(err)
		return
	}
	return
}
