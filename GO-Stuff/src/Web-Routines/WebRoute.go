package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func CheckAndSaveBody(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url,"is DOWN!")
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Response code for",url,"is ==>",resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes,err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fileNameArr := strings.Split(strings.Split(url,"//")[1],".")
				fileName := "Data/"
				fileName += fileNameArr[len(fileNameArr)-2]
				fileName += ".txt"
				err = ioutil.WriteFile(fileName,bodyBytes,0664)
				if err != nil {
					log.Fatal("Error writing file:\n",err)
				}

			} else {
				log.Fatal("Error reading data:\n",err)
			}

		}

	}
}
func main() {
	var urls = []string{"https://golang.org","https://duckduckgo.com","https://google.in","https://en.wikipedia.org/abcd"}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	defer wg.Wait()
	for _,url := range urls {

		go CheckAndSaveBody(url,&wg)

	}
	fmt.Println("No of Go Routines: ",runtime.NumGoroutine())

}
