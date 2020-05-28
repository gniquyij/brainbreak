package main

import (
    "flag"
    "fmt"
    "os/exec"
    "github.com/vjyq/brainbreak/ticker"
)

func main() {
    var countdown int
    flag.IntVar(&countdown, "c", 60, "Define the break. The unit for time is the second. Default is 60s.")
    flag.Parse()
    ticker.Tick(countdown)
    cmd := exec.Command("/usr/bin/say", "Time's up.")
    _, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
}