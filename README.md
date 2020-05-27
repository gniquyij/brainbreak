# brainbreak

Set timer via cmd. 

`$ bb` will:
- start a 60s countdown
- alert when time runs out

You could `$ bb -c <time>` to customize the countdown. 
- e.g. `$ bb -c 300` will set a 5 mins timer

Use `&` to detach the timer. 
- e.g. `$ bb &`

## Setup

```
$ go get github.com/vjyq/brainbreak
$ cd $(go env GOPATH)/src/github.com/vjyq/brainbreak/bin
$ bash shortcut.sh
```
Then just restart your terminal.

## Author

yuqing.ji@outlook.com