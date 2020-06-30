package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	url := os.Getenv("HEALTHCHECKS_URL")
	if url == "" {
		panic("Empty healthcheck url")
	}
	interval := os.Getenv("HEALTHCHECKS_INTERVAL") // add default

	intervalSeconds, err := strconv.Atoi(interval)
	if err != nil {
		panic(fmt.Sprintf("Invalid value for HEALTHCHECKS_INTERVAL: %s", interval))
	}

	for {
		check(url)
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
	}
}

func check(url string) {
	_, err := http.Head(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
