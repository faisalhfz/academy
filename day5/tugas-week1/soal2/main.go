package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func getDescription() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/description")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

func getJobList() string {
	resp, err := http.Get("https://workspace-rho.vercel.app/api/jobs")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body)
}

type Cache struct {
	DescString string
	JobList    string
	IsFilled   bool
}

func (c *Cache) aggregate() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	if !c.IsFilled {

		go func() {
			c.DescString = getDescription()
			wg.Done()
		}()

		go func() {
			c.JobList = getJobList()
			wg.Done()
		}()
		wg.Wait()

		c.IsFilled = true
	}

	fmt.Println(c.DescString)
	fmt.Println(c.JobList)
}

func main() {
	n := 4 // pengulangan

	for i := 1; i <= n; i++ {
		start := time.Now()
		fmt.Println("dari calculate ", start)

		chc := Cache{}
		chc.aggregate()

		fmt.Printf("took %v\n", time.Since(start))

		if i%2 == 0 {
			chc.IsFilled = false
		}
	}
}
