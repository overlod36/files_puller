package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func is_html(content_type string) bool {
	return strings.HasPrefix(content_type, "text/html")
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

func download_file(link string, title string, path string) error {
	outfile, err := os.Create(path + title + ".pdf")
	if err != nil {
		return err
	}
	defer outfile.Close()
	response, err := http.Get(link)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if is_html(response.Header.Get("Content-Type")) {
		fmt.Println("Oh sheesh")
		return nil
	}
	_, err = io.Copy(outfile, response.Body)
	if err != nil {
		return err
	}
	return nil
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
	path_for_download, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path_for_download += "\\result\\"
	for _, journal := range journals {
		for _, article := range journal.Articles {
			err := download_file(article.Link, article.Title, path_for_download)
			if err != nil {
				fmt.Println(err)
			}
		}
		break
	}
}
