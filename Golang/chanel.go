package main

import (
  "fmt"
  "time"
)

func main(){
  t := time.Now()
  fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

  res1 := make(chan int)
  res2 := make(chan int)

  go calculate(10000, res1)
  go calculate(20000, res2)

  fmt.Println(<-res1)
	fmt.Println(<-res2)

	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))
}


func calculate(n int, ch chan int){
  t := time.Now()
  result := 0
  for i := 0; i <= n; i++{
    result += i*2
    time.Sleep(time.Millisecond * 3)
  }
  fmt.Printf("Время выполнения расчетов: %s\n", time.Since(t))
  ch <- result
}
