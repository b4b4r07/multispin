package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX = 3

var repos = []string{
	"b4b4r07/enhancd",
	"b4b4r07/gomi",
	"b4b4r07/dotfiles",
	"zplug/zplug",
	"zsh-users/antigen",
	"fujiwara/nssh",
	"crowi/crowi",
	"mattn/memo",
}

func main() {
	var wg sync.WaitGroup

	q := make(chan string, len(repos))
	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go donwload(&wg, q)
	}

	for i := 0; i < len(repos); i++ {
		q <- repos[i]
	}
	close(q)

	wg.Wait()
}

func donwload(wg *sync.WaitGroup, q chan string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		url, ok := <-q
		if !ok {
			return
		}
		sec := rand.Intn(5)
		fmt.Println(sec, url)
		time.Sleep(time.Duration(sec) * time.Second)
	}
}
