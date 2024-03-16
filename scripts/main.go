package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	now := time.Now()
  opt, _ := redis.ParseURL("rediss://default:a51126d8fd6d4058903fcc0cfd7e856d@eu1-pumped-clam-39849.upstash.io:39849")
  client := redis.NewClient(opt)

	fmt.Println("After connected",time.Since(now))

  client.Set(ctx, "ethan2", "ball", 0)
	fmt.Println("After setted",time.Since(now))
  val := client.Get(ctx, "foo").Val()
	fmt.Println("After getted",time.Since(now))

	client.Set(ctx, "ethan3", "ball", 0)
	fmt.Println("After setted 2",time.Since(now))
  val = client.Get(ctx, "foo").Val()
	fmt.Println("After getted 2",time.Since(now))
  fmt.Println(val)
}