package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSys(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Wait()
		fmt.Println("first")
	}()
	go func() {
		wg.Wait()
		fmt.Println("second")
	}()
	time.Sleep(2 * time.Second)
	wg.Done()
	fmt.Println("end")
}
