package main
import "fmt"
import "math"

func reverse(n int) int {
  rev := 0
  for n != 0 {
    rev = rev*10 + (n%10)
    n /= 10
  }

  return rev
}

func main() {
  n_digit := 3

  lower := int(math.Pow(10,float64(n_digit-1)))
  upper := int(math.Pow(10, float64(n_digit))-1)

  max_prod := 0

  for i := upper; i >= lower; i-- {
    for j := i; j >= lower; j-- {
      prod := i*j
      if prod < max_prod {
        break
      }

      rev := reverse(prod)
      if prod == rev && prod > max_prod {
        max_prod = prod
      }
    }
  }

  fmt.Println(max_prod)
}
