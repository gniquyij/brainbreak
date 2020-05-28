package ticker

import (
    "time"
)

func Tick(countdown int) {
    ticker := time.NewTicker(time.Second)
    for i := 0; i < countdown; i++ {
        <-ticker.C
    }
    ticker.Stop()
}