package main

import (
	// "kartikb385/sample/sec"
	"context"
	"fmt"
)

type APiHandler struct {
}

// func (a *APiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	w.WriteHeader(303)
// 	w.Write([]byte("Hello world"))
// }

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Name", "Kartik")
	receiver(ctx)

}

func receiver(ctx context.Context) {
	fmt.Printf("First Reciver %d", ctx.Value("Name"))
	ctx = context.WithValue(ctx, "Name", "Varun")
	anotherReciver(ctx)
	fmt.Printf("Third Reciver %d", ctx.Value("Name"))
}

func anotherReciver(ctx context.Context) {
	fmt.Printf("Second Reciver %d", ctx.Value("Name"))
}