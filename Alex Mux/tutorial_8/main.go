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
var m = sync.Mutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

// goroutines launch multiple functions and execute it cunrrently

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		/*	the program spawn the task in the background but
			didnt wait until the tasks finish and exit the program before they complete
			Use wait to complete the tasks and then continue, sync.WaitGroup
		*/
		wg.Add(1)
		go dbCall(i)
	}
	/*	wait the counter until the counter goes to 0, which means tasks have completed,
		and the rest of the code executes
	*/
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

// Simulates a call from the db with a delay
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	// add value to the slice
	/*
		but multiple threads modifing the same memory location
		at the same time, can cause memory corruption, unexpected results.
		Solution is to use mutex to control the writting to the slice, to make it safe for the cuncurrent program
	*/

	//	checks if the lock is already been set by another go routine
	// 	if it has, it will wait until the lock is release and set the lock itself
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	// close the waitgroup which decrements the counter
	wg.Done()
}
