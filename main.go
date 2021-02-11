package main

import (
    "fmt"
	"time"

    "github.com/mileusna/crontab"
)

func main() {
    ctab := crontab.New() // create cron table

    ctab.MustAddJob("* * * * *", PrintTimeNow) // every minute
	
	fmt.Println("Crontab Is Running")
    time.Sleep(5 * time.Minute)
}

func PrintTimeNow() {
	layout := "2006-01-02 15:04:05"
    fmt.Println("Time Now:", time.Now().Format(layout))
}
