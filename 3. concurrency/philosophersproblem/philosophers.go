// Implement the dining philosopher’s problem with the following constraints/modifications.

// 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

// 5. The host allows no more than 2 philosophers to eat concurrently.

// 6. Each philosopher is numbered, 1 through 5.

// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) tryLockChopsticks() bool {
	// Randomly decide the order of picking chopsticks
	if rand.Intn(2) == 0 {
		p.leftChopstick.Lock()
		if !p.rightChopstick.TryLock() {
			p.leftChopstick.Unlock()
			time.Sleep(time.Second)
			return p.tryLockChopsticks()
		}
	} else {
		p.rightChopstick.Lock()
		if !p.leftChopstick.TryLock() {
			p.rightChopstick.Unlock()
			time.Sleep(time.Second)
			return p.tryLockChopsticks()
		}
	}
	return true
}

func (p *Philosopher) Eat(wg *sync.WaitGroup, host chan bool) {
	defer wg.Done()
	for i := 0; i < 3; i++ {

		// Lock the philosopher to ensure atomic actions

		go func() {
			defer wg.Done()
			// Request permission from the host to eat
			host <- true

			p.tryLockChopsticks()

			// Critical section: eating
			fmt.Printf("starting to eat %d\n", p.number)
			time.Sleep(time.Second)
			fmt.Printf("finishing eating %d\n", p.number)

			// Unlock the chopsticks
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()

			// Release permission from the host
			<-host
		}()

	}
}

func main() {
	const numPhilosophers = 5
	philosophers := make([]*Philosopher, numPhilosophers)
	chopsticks := make([]*Chopstick, numPhilosophers)
	var wg sync.WaitGroup

	// Create a host channel to allow up to 2 philosophers to eat concurrently
	host := make(chan bool, 2)

	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{}
	}

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
		}
		wg.Add(3)
		go philosophers[i].Eat(&wg, host)
	}

	wg.Wait()
}
