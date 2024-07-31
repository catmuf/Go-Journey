package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// mutual exclusion
// it completely locks out other go routine to accessing to result slice
// go provides other mutex called RWMutex{}, same functionalities as Mutex but adds RLock and RUnlock
var m = sync.RWMutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// goroutines launch multiple functions and execute it cunrrently

func main() {
	t0 := time.Now()
	for i := 0; i < 4; i++ {
		/*	the program spawn the task in the background but
			didnt wait until the tasks finish it and closed the program.
			Use wait to complete the tasks and then continue, sync.WaitGroup
		*/
		wg.Add(1)
		go count()
	}
	/*	wait the counter until the counter goes to 0, which means tasks have completed,
		and the rest of the code executes
	*/
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func count() {
	var res int
	for i := 0; i < 1000000000; i++ {
		res += 1
	}
	fmt.Println(res)
	wg.Done()
}
