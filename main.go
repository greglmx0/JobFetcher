package main

import (
    "fmt"
    "github.com/robfig/cron"
)

func main() {
    c := cron.New()
    c.AddFunc("*/5 * * * * *", func() {
        fmt.Println("Hello, world! 2")
    })
    c.Start()

    select {}
}
