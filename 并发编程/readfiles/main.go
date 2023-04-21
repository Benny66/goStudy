package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 编写一个程序，使用goroutines并发地处理文件，并对每个文件计算哈希值。
func hashFile(path string, result chan<- string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		result <- ""
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Error while hashing", path, "-", err)
		result <- ""
		return
	}

	result <- fmt.Sprintf("%x", hash.Sum(nil))
}

func walkFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func main() {
	root := "/Users/shahao/Project/goStudy/并发编程"

	paths, err := walkFiles(root)
	if err != nil {
		panic(err)
	}

	results := make(chan string)

	for _, path := range paths {
		go hashFile(path, results)
	}

	for i := 0; i < len(paths); i++ {
		fmt.Printf("%s: %s\n", paths[i], <-results)
	}
}
