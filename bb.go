package main

import (
    "flag"
    "fmt"
    "os/exec"
    "time"
)

func main() {
    var countdown int
    flag.IntVar(&countdown, "c", 5, "Define the break. Default is 5s.")
    flag.Parse()
    ticker := time.NewTicker(time.Second)
    for i := 0; i < countdown; i++ {
        <-ticker.C
    }
    ticker.Stop()
    cmd := exec.Command("/usr/bin/say", "Time's up.")
    _, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
}