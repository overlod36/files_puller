package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Journal struct {
	Link     string    `json:"link"`
	Title    string    `json:"title"`
	Decade   string    `json:"decade"`
	Year     string    `json:"year"`
	Articles []Article `json:"articles"`
}

type Article struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func get_parse_files_slice() ([]string, error) {
	var files []string
	root, err := os.Getwd()
	root += "\\files\\"
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			files = append(files, strings.ReplaceAll(path, root, ""))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, err
}

func open_file(filename string) ([]Journal, error) {
	var journals []Journal
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(root + "\\files\\" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byte_value, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byte_value, &journals)
	return journals, err
}

func main() {
	parse_files, err := get_parse_files_slice()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("[Результаты парсинга]")
	for _, file := range parse_files {
		fmt.Println(file)
	}
	fmt.Print(">> ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	filename = strings.TrimSuffix(filename, "\n")
	filename = strings.TrimSuffix(filename, "\r")
	if err != nil {
		fmt.Println(err)
	}
	journals, err := open_file(filename)
	if err != nil {
		fmt.Println(err)
	}
	for _, journal := range journals {
		fmt.Printf("%+v\n", journal)
		fmt.Println()
	}
}
