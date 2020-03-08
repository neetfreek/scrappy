package lib

import (
	"fmt"
	"sync"
	"time"
)

// test slices of strings
var cats = []string{"cheshire", "ginger", "salmon", "longhair", "persian", "siamese", "tabby"}
var dogs = []string{"beagle", "salmon", "poodle", "longhair", "spaniel"}
var fish = []string{"perch", "salmon", "longhair", "trout"}
var animals = [][]string{cats, dogs, fish}
var all = []string{}

// Timer prints how long a routine takes
func Timer(start time.Time, message string) {
	time := time.Since(start)
	fmt.Printf("%v: %v", message, time)
}

// Test concurrency
func Test() {
	defer Timer(time.Now(), "Operation time")

	getAllAnimals()
}

func getAllAnimals() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for _, item := range animals {
		wg.Add(1)
		go sendToAll(item, &wg, &m)
	}
	wg.Wait()
}

func sendToAll(animals []string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	for _, item := range animals {
		if !itemInSlice(item, all) {
			m.Lock()
			all = append(all, item)
			m.Unlock()
		} else {
		}
	}
}
