package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("Starting")
  win := float64(0)
  run := float64(1000000)
  random := rand.Intn(3)
  fmt.Println("Random Chosen is:", random)
  for j := 0; j < 1000000; j++ {
    firstChoice := rand.Intn(3) //Initial Guess
    prize := rand.Intn(3) //Prize set
    open := newChoice(firstChoice, prize) //open non prize or guess
    swap := newChoice(firstChoice, open) //Swap to non opened
    if swap == prize {
      win++
    } else {
      //fmt.Println("Loose")
    }
  }
  fmt.Println("Win:", win)
  fmt.Println("Win percentage:", win / run * 100, "%")
}

func newChoice(input int, prize int) int {
  if input != 0 && prize != 0 {
    return 0
  }
  if input != 1 && prize != 1 {
    return 1
  }
  if input != 2 && prize != 2 {
    return 2
  }
  return -1;
}
