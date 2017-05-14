package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/morikuni/aec"
)

var spinners = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// const items = 10

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
	items := len(repos)
	fmt.Print(aec.Hide)
	defer func() {
		fmt.Print(aec.Show)
	}()

	for i := 0; i < items; i++ {
		fmt.Printf("\n")
	}
	up := aec.Up(uint(items))
	cmax := len(spinners)
	c := 0
	rand.Seed(time.Now().UnixNano())

	task := make(chan string)
	taskquit := make(chan bool)
	workerquit := make(chan bool)

	// go func() {
	// loop:
	// 	for {
	// 		select {
	// 		case <-taskquit:
	// 			workerquit <- true
	// 			break loop
	// 		case job := <-task:
	// 			_ = job
	// 			// fmt.Println(job)
	// 		}
	// 	}
	// }()

	go func() {
		for _, repo := range repos {
			task <- repo
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
		taskquit <- true
	}()

	for {
		fmt.Print(up)
		if c >= cmax {
			c = 0
		}
		for _, repo := range repos {
			fmt.Print(aec.EraseLine(2))
			if false {
				fmt.Println("✔", "donwloaded", repo)
			} else {
				fmt.Println(spinners[c], "donwloading", repo)
			}
		}
		c++
		time.Sleep(time.Duration(100) * time.Millisecond)
	}

	<-workerquit
}
