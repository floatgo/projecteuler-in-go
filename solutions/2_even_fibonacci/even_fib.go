package main
import "fmt"

func main() {
  sum := 0
  
  a := 0
  b := 1
  c := 0
  for {
    c = a + b

    if c > 4000000 {
      break
    }
    if c % 2 == 0 {
      sum += c
    }
    a = b
    b = c
  }
  fmt.Println(sum)
}
