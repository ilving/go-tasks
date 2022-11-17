package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	time.Sleep(1 * time.Millisecond)
	doDbRequest(ctx)
}

func doDbRequest(ctx context.Context) {
	newCtx, _ := context.WithTimeout(ctx, 8*time.Second)
	timer := time.NewTimer(10 * time.Second)
	select {
	case <-newCtx.Done():
		fmt.Println("Timeout")
	case <-timer.C:
		fmt.Println("Request Done")
	}
}
