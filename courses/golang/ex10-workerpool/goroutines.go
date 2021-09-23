package goroutines

import (
	"bufio"
	"fmt"
	// "net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Job struct {
	id    int
	sleep float64
}

func worker(jobs chan Job, wg *sync.WaitGroup) {
	id := 0
	for job := range jobs {
		if id == 0 {
			id = job.id
			fmt.Printf("worker:%d spawning\n", id)
			defer fmt.Printf("worker:%d stopping\n", id)
		}
		fmt.Printf("worker:%d sleep:%.1f\n", id, job.sleep)
		time.Sleep(time.Duration(job.sleep*1000) * time.Millisecond)
	}
	wg.Done()
}

func getJobs(jobs chan Job) {
	id := 1
	for {
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if s == "" {
			break
		}

		sleep, err := strconv.ParseFloat(strings.Trim(s, "\n"), 64)
		if err != nil {
			fmt.Println(err)
			continue
		}

		jobs <- Job{
			id:    id,
			sleep: sleep,
		}
		id++
	}
	close(jobs)
}

func Run(poolSize int) {
	jobs := make(chan Job, 10)
	go getJobs(jobs)

	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	/*
	   http.HandleFunc("/stats", func(w http.ResponseWriter, req *http.Request) {
	       fmt.Fprintf(w, "%d", len(jobs))
	   })
	   http.ListenAndServe(":8090", nil)
	*/
	wg.Wait()
}
