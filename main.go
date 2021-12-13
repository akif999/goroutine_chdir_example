package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baseDir, _ := os.Getwd()
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go doSomeThingWithChdir(baseDir, i)
	}
	wg.Wait()
}

func doSomeThingWithChdir(baseDir string, num int) {
	defer wg.Done()
	wd, _ := os.Getwd()
	fmt.Printf("routine %v: starting in dir: %s\n", num, wd)
	os.Chdir(filepath.Join(baseDir, fmt.Sprintf("%v", num)))
	wd, _ = os.Getwd()
	fmt.Printf("routine %v: going to sleep in dir: %s\n", num, wd)
	time.Sleep(1000 * time.Millisecond)
	wd, _ = os.Getwd()
	fmt.Printf("routine %v: woke up in dir: %s\n", num, wd)
}
