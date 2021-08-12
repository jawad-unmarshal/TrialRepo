package main

import (
	"fmt"
	"io/ioutil"
	"time"

	// "log"
	"net/http"
	// "runtime"
	"strings"
	// "sync"
)

func CheckAndSaveBody(url string, c chan string) {
	// defer wg.Done()
	resp, err := http.Get(url)
	var s string
	if err != nil {
		s += fmt.Sprintln(url, "is DOWN!")
		s += fmt.Sprintln(err)
		//  c <- s
	} else {
		defer resp.Body.Close()
		s += fmt.Sprintln("Response code for", url, "is ==>", resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fileNameArr := strings.Split(strings.Split(url, "//")[1], ".")
				fileName := "Data/"
				fileName += fileNameArr[len(fileNameArr)-2]
				fileName += ".txt"
				err = ioutil.WriteFile(fileName, bodyBytes, 0664)
				if err != nil {
					s += fmt.Sprintln("Error writing file:\n", err)
				} else {
					s += fmt.Sprintln("Data for ==> ", url, "successfully written")
				}

			} else {
				s += fmt.Sprintln("Error reading data:\n", err)
			}

		}
		// c <- s

	}
	c <- s
}

func ChkAlive(url string, c chan string) {
	// defer wg.Done()
	resp, err := http.Get(url)
	var s string
	if err != nil {
		s += fmt.Sprintln(url, "is DOWN!")
		s += fmt.Sprintln(err)
		c <- url
	} else {
		defer resp.Body.Close()
		s += fmt.Sprintln("Response code for", url, "is ==>", resp.StatusCode)
		c <- url
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fileNameArr := strings.Split(strings.Split(url, "//")[1], ".")
				fileName := "Data/"
				fileName += fileNameArr[len(fileNameArr)-2]
				fileName += ".txt"
				err = ioutil.WriteFile(fileName, bodyBytes, 0664)
				if err != nil {
					s += fmt.Sprintln("Error writing file:\n", err)
				} else {
					s += fmt.Sprintln("Data for ==> ", url, "successfully written")
				}

			} else {
				s += fmt.Sprintln("Error reading data:\n", err)
			}

		}
		// c <- s
		fmt.Println(s)

	}
}

func main() {
	var urls = []string{"https://golang.org", "https://duckduckgo.com", "https://google.in", "https://en.wikipedia.org/abcd"}
	// var wg sync.WaitGroup

	// wg.Add(len(urls))
	// defer wg.Wait()
	var OpChan = make(chan string)
	for _, url := range urls {

		go ChkAlive(url, OpChan)

	}
	// fmt.Println("No of Go Routines: ", runtime.NumGoroutine())
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-OpChan)
	// }

	// for {
	// 	time.Sleep(time.Second * 2)
	// 	go ChkAlive(<-OpChan, OpChan)
	// 	fmt.Println(strings.Repeat("#-##-#", 10))
	// }

	// for url := range OpChan {
	// 	time.Sleep(time.Second * 2)
	// 	go ChkAlive(url, OpChan)
	// 	fmt.Println(strings.Repeat("#-##-#", 10))
	// }

	for url := range OpChan {
		go func(u string) {
			time.Sleep(time.Second * 2)
			ChkAlive(u, OpChan)
			fmt.Println(strings.Repeat("#-##-#", 10))
		}(url)
	}

}
