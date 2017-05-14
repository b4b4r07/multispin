package main

import (
	"fmt"
	"math/rand"
	"time"
)

var repos = []string{
	"b4b4r07/enhancd",
	"b4b4r07/gomi",
	"b4b4r07/dotfiles",
	"zplug/zplug",
	"zsh-users/antigen",
	"fujiwara/nssh",
}

func heavyTask(sec int) {
	fmt.Println(sec)
	time.Sleep(time.Duration(sec) * time.Second)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(repos); i++ {
		heavyTask(rand.Intn(5))
	}
}
