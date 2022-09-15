package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	object struct {
		Player string
		Loop   int
		Value  int
		Err    error
	}
)

func main() {
	var c = make(chan *object)
	defer fmt.Println("GAME OVER")
	for {
		time.Sleep(1 * time.Second)
		go play("Player 1", c)
		go play("Player 2", c)
		go play("Player 3", c)
		go play("Player 4", c)

		tmp := <-c
		fmt.Println(tmp.Player, " Get ", tmp.Value)

		if tmp.Err != nil {
			fmt.Println(tmp.Err.Error())
			break
		}
	}

}

func play(name string, c chan *object) {
	random := rand.Intn(100-1) + 1
	var err error
	if (random % 11) == 0 {
		err = fmt.Errorf("%s Defeat ", name)
	}
	c <- &object{
		Player: name,
		Loop:   1,
		Value:  random,
		Err:    err,
	}

}
