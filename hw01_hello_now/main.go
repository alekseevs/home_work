package main

import (
	"fmt"
	"log"
	"github.com/beevik/ntp"
	"time"
	"os"
)
// current time: <время>
// exact time: <время>
// (без учёта монотонного времени (без "m=±<value>" в выводе))

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05 -0700 UTC")
}

func main() {
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	localTime := time.Now()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("current time: %s\n", formatTime(localTime))
	fmt.Printf("exact time: %s\n", formatTime(exactTime))
}