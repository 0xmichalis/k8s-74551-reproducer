package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	timeout  = flag.Duration("timeout", 1*time.Minute, "Time to send requests before exiting")
	interval = flag.Duration("interval", 1*time.Second, "Interval between sending two requests")
)

func doRequest(client *http.Client, method string, data []byte) error {
	req, err := http.NewRequest(method, "http://localhost:8080/foo", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	return nil
}

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	exit := time.After(*timeout)
	try := time.Tick(*interval)
	for {
		select {
		case <-exit:
			os.Exit(0)

		case <-try:
			// generate random size data when doing a POST request
			random := r.Intn(5000000)
			method := http.MethodPost
			var data []byte

			if (random % 5) == 0 {
				method = http.MethodGet
			}
			logLine := fmt.Sprintf("Sending request %s", method)
			if method == http.MethodPost {
				data = bytes.Repeat([]byte("r"), random)
				logLine += fmt.Sprintf(" with size %d", random)
			}
			log.Print(logLine)
			if err := doRequest(client, method, data); err != nil {
				log.Print(err)
			}
		}
	}
}
