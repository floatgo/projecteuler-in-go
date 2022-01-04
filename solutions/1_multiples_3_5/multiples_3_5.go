package main
import "fmt"

func main() {
  num := 1000
  sm := 0
  for i := 1; i < num; i ++ {
    if i % 3 == 0 || i % 5 == 0 {
      sm += i
    }
  }
  fmt.Println(sm)
}
