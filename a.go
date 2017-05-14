package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/morikuni/aec"
)

var spinners = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

const items = 10

func main() {
	fmt.Print(aec.Hide)
	defer func() {
		fmt.Print(aec.Show)
	}()
	for i := 1; i <= items; i++ {
		fmt.Printf("\n")
	}
	up := aec.Up(items)

	cmax := len(spinners)
	c := 0
	done := false
	rand.Seed(time.Now().UnixNano())

	for {
		if rand.Intn(50) == 0 {
			done = true
		}
		fmt.Print(up)
		if c >= cmax {
			c = 0
		}
		for i := 1; i <= items; i++ {
			fmt.Print(aec.EraseLine(2))
			if done {
				fmt.Println("✔", "donwloaded")
			} else {
				fmt.Println(spinners[c], "donwloading")
			}
		}
		c++
		time.Sleep(time.Duration(100) * time.Millisecond)
		if done {
			break
		}
	}
}
