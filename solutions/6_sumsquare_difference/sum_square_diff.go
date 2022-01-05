package main
import "fmt"

func main() {
  n := 100

  sum_squares_first_n := int(n*(n+1)*(2*n + 1)/6)
  sum_first_n := int(n*(n+1)/2)

  fmt.Println((sum_first_n*sum_first_n) - sum_squares_first_n)
}
