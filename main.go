package main

import (
  "github.com/tulios/load_balancer/load_balancer"
	"os"
  "net/http"
  "fmt"
)

func main() {
	balancer, err := load_balancer.New(os.Args[1:]...)
  if err != nil { panic(err) }

  http.Handle("/", balancer)
  fmt.Println("Go LoadBalancer listening 0.0.0.0:8081")
  http.ListenAndServe(":8081", nil)
}
