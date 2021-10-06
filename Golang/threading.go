package main

import (
    "fmt"
    //"time"
)


/*
func main(){
  c := make(chan int)
  for i := 0; i < 5; i++{
    go sleepGopher(i, c)
  }
  for i := 0; i < 5; i++{
    gopherID := <- c
    fmt.Println("gopher ", gopherID, " has finished sleeping")
  }
}

func sleepGopher(id int, c chan int){
  time.Sleep(3 * time.Second)
  fmt.Println("... ", id, " snore ...")
  c <- id
}
*/

/*
func main(){
  for i := 1; i <= 10; i++{
    go factorial(i)
  }
  fmt.Scanln()
  fmt.Println("The end")
}

func factorial(n int){
  if (n < 1){
    fmt.Println("The end")
    return
  }
  result := 1
  for i := 1; i <= n; i++{
    result *= i
  }
  fmt.Println(n , "-", result)
}
*/

func main(){
  intCh := make(chan int)
  for i := 1; i <= 5; i++{
    go factorial(i, intCh)
    fmt.Println(<-intCh)
  }

  fmt.Println("The End")

}
func factorial(n int, ch chan int){
  result := 1
  for i := 1; i <= n; i++{
    result *= i
  }
  fmt.Println(n, "-", result)
  ch <- result
}
