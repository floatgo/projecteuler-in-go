package main
import "fmt"

func main() {
  c := 0

  sum := 1000
  for a := 1; a < sum; a++ {
    for b := 2; b < sum; b++  {
      c = (sum - a - b)
      if c <= 0 {
        continue
      }
      if (a*a) + (b*b) == (c*c) {
        fmt.Println(a*b*c)
      }
    }
  }
}
