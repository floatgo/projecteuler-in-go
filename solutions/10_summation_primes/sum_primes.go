package main
import "fmt"

func main() {
  max := 2000000
  
  sieve := make([]bool, max)
  for i := 0; i < max; i++ {
    sieve[i] = true
  }

  sieve[1] = false

  for i := 2; i < max; i++ {
    if sieve[i] == true {
      for j := 2*i; j < max; j += i {
        sieve[j] = false
      }
    }
  }

  sum := 0
  for i := 2; i < max; i++ {
    if sieve[i] == true {
      sum += i
    }
  }

  fmt.Println(sum)
}
