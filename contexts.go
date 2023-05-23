package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func entry_5() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	//time.Sleep(time.Second * 2)
	cancel()
	wg.Wait()
	for i := 0; i < 100; i++ {
		if i == 10 {
			break
		}
		fmt.Println("print_tt i ", i)
	}
	fmt.Println("main close")
}

func worker(ctx context.Context) {
	//go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker on")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2 on")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker2 for finish")
	wg.Done()
}
