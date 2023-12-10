package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Article struct {
	title string
	link  string
}

type Journal struct {
	title    string
	link     string
	articles []Article
	decade   int64
	year     int64
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

func open_file(filename string) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	file, err := os.Open(root + "\\files\\" + filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
	err = open_file(filename)
	if err != nil {
		fmt.Println(err)
	}
}
