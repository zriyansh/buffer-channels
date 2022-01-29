// to represent race condition, when stingy adds 10 and spendy substracts 10
// at the same time

package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("spendy Done")
}

func main() {
	// both are initialised as threads, as 'go' is mentioned
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}

// this program should give wrong output (not 100), but idk why it's not
// race condition should occur, but it's working fine.
