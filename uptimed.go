package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func AppCleanup() {
	log.Println("CLEANUP APP BEFORE EXIT!!!")
}

var (
	url = flag.String(
		"url",
		os.Getenv("UPTIMED_URL"),
		"Set the UPTIMED_URL env var or pass the url of the request that uptimed should do occasionally",
	)
	interval = flag.Int("interval", 30, "Set the polling interval")
)

func main() {

	flag.Parse()

	sleepDuration := time.Duration(*interval)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGQUIT)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		s := <-sigs
		s.Signal()
		log.Printf("RECEIVED SIGNAL: %s", s)
		AppCleanup()
		os.Exit(1)
	}()

	for {

		req, err := http.NewRequest("GET", *url, nil)

		if err != nil {
			log.Println("Not reachable, sleeping...")
			time.Sleep(sleepDuration * time.Second)
			panic("Bailing out")
		}
		log.Println("Reached mothership.")

		_, err = http.DefaultClient.Do(req)

		if err != nil {
			log.Println("Not reachable, sleeping...")
			time.Sleep(sleepDuration * time.Second)
			panic("Bailing out")
		}

		time.Sleep(sleepDuration * time.Second)

	}
}
