package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Website struct {
	Name string
	URL  string
}

// 编写一个程序，使用goroutine实现并发爬取多个网站的数据，并将结果保存到数据库中。
func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/webdata")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	websites := []Website{
		{Name: "Google", URL: "https://www.google.com"},
		{Name: "Facebook", URL: "https://www.facebook.com"},
		{Name: "Twitter", URL: "https://www.twitter.com"},
		{Name: "Amazon", URL: "https://www.amazon.com"},
	}

	insertStmt, err := db.Prepare("INSERT INTO website_data (name, data) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertStmt.Close()

	maxMemory := int64(64 * 1024 * 1024)
	freeMemory := maxMemory
	semaphore := make(chan struct{}, 2)
	for _, w := range websites {
		semaphore <- struct{}{}
		go func(w Website) {
			defer func() { <-semaphore }()

			client := &http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost: 1,
				},
			}
			req, err := http.NewRequest("GET", w.URL, nil)
			if err != nil {
				fmt.Printf("%s: %s\n", w.Name, err.Error())
				return
			}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("%s: %s\n", w.Name, err.Error())
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s: %s\n", w.Name, err.Error())
				return
			}

			bodySize := int64(len(body))
			if bodySize > freeMemory {
				fmt.Printf("%s: downloaded data is too large (%d bytes)\n", w.Name, bodySize)
				return
			}

			freeMemory -= bodySize

			_, err = insertStmt.Exec(w.Name, body)
			if err != nil {
				fmt.Printf("%s: %s\n", w.Name, err.Error())
				return
			}

			fmt.Printf("%s: %d bytes downloaded\n", w.Name, bodySize)

			freeMemory += bodySize
		}(w)
	}

	// 等待所有goroutine完成
	for i := 0; i < cap(semaphore); i++ {
		semaphore <- struct{}{}
	}
	time.Sleep(time.Second * 10)
}
