package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2) // telling to use 2 processors

	var waitGrp sync.WaitGroup
	waitGrp.Add(2) // number indicate number of functions to wait for

	// Go concurrecny model is Actor Model or Communicating sequential Process(CSP)
	// "go" keyword use to make a function goroutine
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Sumair")
	}()

	waitGrp.Wait()

}
