// to represent race condition, when stingy adds 10 and spendy substracts 10
// at the same time

package s

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 10
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("stingy sees balance of ", money)
		moneyDeposited.Signal() // we signal that we deposited the money, spendy thread will wake up, try to acquire lock again
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait() // wait till we have more than 20 to spend and releases the lock if condition is not met.
		}
		money -= 20
		fmt.Println("spendy sees balance of ", money)
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
