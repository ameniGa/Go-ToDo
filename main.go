package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type I interface {
   set(int) string
}
type S struct {
I
}

func myfunc(nb int,i I) {
	fmt.Println(i.set(nb))
}

func main() {
	var s int
	myfunc(5,s)


}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		log.Println(s)
	case <-ctx.Done():
		log.Printf("done %v", ctx.Err())
	}
}
