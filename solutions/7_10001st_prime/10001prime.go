package main
import "fmt"

func main() {
  max := 1000000
  
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

  count := 0
  for i := 2; i < max; i++ {
    if sieve[i] == true {
      count += 1
    }

    if count == 10001 {
      fmt.Println(i)
      break
    }
  }
}
